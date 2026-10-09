"""Run disposable native-certificate Terraform UAT; credentials stay in environment."""
import argparse
import hashlib
import json
import os
from pathlib import Path
import subprocess
import tempfile
import urllib.request
import urllib.error

parser = argparse.ArgumentParser()
parser.add_argument("--provider-dir", required=True)
parser.add_argument("--namespace", required=True)
parser.add_argument("--name", required=True)
parser.add_argument("--certificate", required=True)
parser.add_argument("--private-key", required=True)
parser.add_argument("--receipt", required=True)
args = parser.parse_args()
if not os.environ.get("XCSH_API_URL") or not os.environ.get("XCSH_API_TOKEN"):
    raise SystemExit("Tenant credentials are required in the environment")
receipt = {"version": 1, "gates": {}, "cleanup": False}
with tempfile.TemporaryDirectory(prefix="blindfold-uat-") as temp:
    root = Path(temp)
    root.chmod(0o700)
    (root / "dev.tfrc").write_text("provider_installation {\n dev_overrides {\n" + json.dumps("f5-sales-demo/xcsh") + " = " + json.dumps(str(Path(args.provider_dir).resolve())) + "\n }\n direct {}\n}\n")
    config = 'terraform {\n required_providers { xcsh = { source = "f5-sales-demo/xcsh" } }\n}\nprovider "xcsh" {}\nresource "xcsh_certificate" "uat" {\n'
    for key, value in {"name": args.name, "namespace": args.namespace}.items():
        config += key + " = " + json.dumps(value) + "\n"
    config += "blindfold = {\n certificate_file = " + json.dumps(str(Path(args.certificate).resolve())) + "\n private_key_file = " + json.dumps(str(Path(args.private_key).resolve())) + "\n}\n}\n"
    (root / "main.tf").write_text(config)
    env = dict(os.environ, TF_CLI_CONFIG_FILE=str(root / "dev.tfrc"), TF_IN_AUTOMATION="1")
    def run(stage, *argv, allowed=(0,)):
        completed = subprocess.run(["terraform", *argv], cwd=root, env=env, capture_output=True, check=False)
        (root / (stage + ".log")).write_bytes(completed.stdout + completed.stderr)
        receipt["gates"][stage] = completed.returncode
        if completed.returncode not in allowed:
            raise RuntimeError(stage + " failed; private log retained until cleanup")
    try:
        run("init", "init", "-no-color")
        run("plan", "plan", "-no-color", "-out=initial.plan")
        run("apply", "apply", "-no-color", "initial.plan")
        run("empty_plan", "plan", "-no-color", "-detailed-exitcode")
        private = Path(args.private_key).read_bytes()
        state = (root / "terraform.tfstate").read_bytes()
        if private in state or private.splitlines()[1] in state:
            raise RuntimeError("plaintext private input found in state")
        receipt["gates"]["secret_absence"] = True
        receipt["state_sha256"] = hashlib.sha256(state).hexdigest()
    finally:
        run("destroy_plan", "plan", "-destroy", "-out=destroy.plan", "-no-color")
        run("destroy", "apply", "-no-color", "destroy.plan")
        url = os.environ["XCSH_API_URL"].rstrip("/") + "/api/config/namespaces/" + args.namespace + "/certificates/" + args.name
        request = urllib.request.Request(url, headers={"Authorization": "APIToken " + os.environ["XCSH_API_TOKEN"]})
        try:
            urllib.request.urlopen(request, timeout=30)
        except urllib.error.HTTPError as error:
            if error.code == 404:
                receipt["cleanup"] = True
            else:
                raise
        if not receipt["cleanup"]:
            raise RuntimeError("cleanup readback failed")
Path(args.receipt).write_text(json.dumps(receipt, indent=2) + "\n")
