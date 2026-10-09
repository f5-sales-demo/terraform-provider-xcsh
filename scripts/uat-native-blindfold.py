# Executable command filename follows existing script naming.
# pylint: disable=invalid-name
"""Run disposable native-certificate Terraform UAT; credentials stay in environment."""

import argparse
import hashlib
import json
import os
import re
import shutil
import subprocess
import tempfile
import urllib.error
import urllib.parse
import urllib.request
from pathlib import Path
from typing import TypedDict

NOT_FOUND = 404
LABEL = re.compile(r"^[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?$")
parser = argparse.ArgumentParser()
parser.add_argument("--provider-dir", required=True)
parser.add_argument("--namespace", required=True)
parser.add_argument("--name", required=True)
parser.add_argument("--certificate", required=True)
parser.add_argument("--private-key", required=True)
parser.add_argument("--receipt", required=True)
args = parser.parse_args()
api_url = os.environ.get("XCSH_API_URL", "").rstrip("/")
origin = urllib.parse.urlparse(api_url)
invalid_origin = origin.scheme != "https" or bool(origin.username or origin.password)
invalid_origin = invalid_origin or bool(origin.query or origin.fragment or origin.path)
if invalid_origin:
    message = "Tenant API URL must be an HTTPS origin"
    raise SystemExit(message)
if (
    not os.environ.get("XCSH_API_TOKEN")
    or not LABEL.fullmatch(args.namespace)
    or not LABEL.fullmatch(args.name)
):
    message = "Tenant credentials and valid disposable identities are required"
    raise SystemExit(message)
terraform_path = shutil.which("terraform")
if not terraform_path:
    message = "Terraform executable is required"
    raise SystemExit(message)
terraform: str = terraform_path


class Receipt(TypedDict, total=False):
    """Sanitized acceptance facts without tenant credentials."""

    version: int
    gates: dict[str, int | bool]
    cleanup: bool
    state_sha256: str


receipt: Receipt = {"version": 1, "gates": {}, "cleanup": False}
with tempfile.TemporaryDirectory(prefix="blindfold-uat-") as temp:
    root = Path(temp)
    root.chmod(0o700)
    provider_path = json.dumps(str(Path(args.provider_dir).resolve()))
    (root / "dev.tfrc").write_text(
        f'provider_installation {{\n dev_overrides {{\n "f5-sales-demo/xcsh" = {provider_path}\n }}\n direct {{}}\n}}\n'
    )
    config = 'terraform {\n required_providers { xcsh = { source = "f5-sales-demo/xcsh" } }\n}\nprovider "xcsh" {}\nresource "xcsh_certificate" "uat" {\n'
    for key, value in {"name": args.name, "namespace": args.namespace}.items():
        config += key + " = " + json.dumps(value) + "\n"
    config += (
        "blindfold = {\n certificate_file = "
        + json.dumps(str(Path(args.certificate).resolve()))
        + "\n private_key_file = "
        + json.dumps(str(Path(args.private_key).resolve()))
        + "\n}\n}\n"
    )
    (root / "main.tf").write_text(config)
    env = dict(
        os.environ, TF_CLI_CONFIG_FILE=str(root / "dev.tfrc"), TF_IN_AUTOMATION="1"
    )

    def run(stage: str, *argv: str) -> None:
        """Run fixed Terraform operations and retain only private temporary logs."""
        # Terraform receives fixed command tokens and validated paths, without a shell.
        completed = subprocess.run(  # noqa: S603 -- fixed Terraform tokens and validated paths
            [terraform, *argv], cwd=root, env=env, capture_output=True, check=False
        )
        (root / (stage + ".log")).write_bytes(completed.stdout + completed.stderr)
        receipt["gates"][stage] = completed.returncode
        if completed.returncode:
            message = stage + " failed; private log retained until cleanup"
            raise RuntimeError(message)

    try:
        run("init", "init", "-no-color")
        run("plan", "plan", "-no-color", "-out=initial.plan")
        run("apply", "apply", "-no-color", "initial.plan")
        run("empty_plan", "plan", "-no-color", "-detailed-exitcode")
        private = Path(args.private_key).read_bytes()
        state = (root / "terraform.tfstate").read_bytes()
        if private in state or private.splitlines()[1] in state:
            message = "plaintext private input found in state"
            raise RuntimeError(message)
        receipt["gates"]["secret_absence"] = True
        receipt["state_sha256"] = hashlib.sha256(state).hexdigest()
    finally:
        run("destroy_plan", "plan", "-destroy", "-out=destroy.plan", "-no-color")
        run("destroy", "apply", "-no-color", "destroy.plan")
        url = (
            api_url
            + "/api/config/namespaces/"
            + args.namespace
            + "/certificates/"
            + args.name
        )
        request = urllib.request.Request(  # noqa: S310 -- validated HTTPS origin
            url, headers={"Authorization": "APIToken " + os.environ["XCSH_API_TOKEN"]}
        )
        try:
            with urllib.request.urlopen(request, timeout=30) as response:  # noqa: S310 -- validated HTTPS origin
                response.read(1)
        except urllib.error.HTTPError as error:
            if error.code == NOT_FOUND:
                receipt["cleanup"] = True
            else:
                raise
        if not receipt["cleanup"]:
            message = "cleanup readback failed"
            raise RuntimeError(message)
Path(args.receipt).write_text(json.dumps(receipt, indent=2) + "\n")
