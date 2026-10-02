"""Exercise HTTP load-balancer lifecycle using Terraform and a loopback-only API."""

import argparse
import json
import os
import subprocess
import tempfile
import threading
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path
from typing import Any


def require(condition: bool, message: object) -> None:
    """Raise on acceptance failure even when Python optimization is enabled."""
    if not condition:
        raise AssertionError(str(message))


class API:
    """Synthetic XC API: type-changing PUT is forbidden and POST creates a new UID."""

    def __init__(self) -> None:
        """Initialize an empty synthetic API."""
        self.object: dict[str, Any] | None = None
        self.sequence = 0
        self.writes: list[str] = []
        self.type_changing_puts = 0

    def current(self) -> dict[str, Any]:
        """Return the current object or fail when a write lost it."""
        if self.object is None:
            message = "synthetic object is absent"
            raise RuntimeError(message)
        return self.object

    @staticmethod
    def selection(spec: dict[str, Any]) -> str:
        """Return the selected type without inferring defaults."""
        return next(
            (name for name in ("http", "https", "https_auto_cert") if name in spec),
            "omitted",
        )

    def handler(self) -> type[BaseHTTPRequestHandler]:
        """Build an HTTP handler for the synthetic object."""
        api = self

        class Handler(BaseHTTPRequestHandler):
            """Serve only the synthetic resource lifecycle."""

            def log_message(self, _format: str, *args: object) -> None:  # pylint: disable=arguments-differ
                """Suppress synthetic HTTP request logging."""

            def reply(self, status: int, body: object) -> None:
                """Encode a synthetic JSON response."""
                self.send_response(status)
                self.send_header("Content-Type", "application/json")
                self.end_headers()
                self.wfile.write(json.dumps(body).encode())

            def do_GET(self) -> None:  # pylint: disable=invalid-name
                """Read the synthetic object."""
                if api.object is None:
                    self.reply(404, {"message": "synthetic object absent"})
                else:
                    self.reply(200, api.object)

            def do_POST(self) -> None:  # pylint: disable=invalid-name
                """Create a new synthetic identity."""
                if api.object is not None:
                    self.reply(409, {"message": "synthetic name already exists"})
                    return
                body = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
                api.sequence += 1
                body["metadata"]["uid"] = f"synthetic-object-{api.sequence}"
                body["resource_version"] = str(api.sequence)
                api.object = body
                api.writes.append("POST")
                self.reply(200, body)

            def do_PUT(self) -> None:  # pylint: disable=invalid-name
                """Update settings while enforcing immutable selection."""
                body = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
                if api.selection(body["spec"]) != api.selection(api.current()["spec"]):
                    api.type_changing_puts += 1
                    self.reply(400, {"message": "loadbalancer_type is immutable"})
                    return
                body["metadata"]["uid"] = api.current()["metadata"]["uid"]
                body["resource_version"] = str(api.sequence)
                api.object = body
                api.writes.append("PUT")
                self.reply(200, body)

            def do_DELETE(self) -> None:  # pylint: disable=invalid-name
                """Delete the synthetic object."""
                api.object = None
                api.writes.append("DELETE")
                self.reply(200, {})

        return Handler


# This CLI orchestrates independent acceptance scenarios against one isolated runner.
# pylint: disable-next=too-many-locals,too-many-statements
def exercise(args: argparse.Namespace) -> None:
    """Verify Terraform plans and writes using a disposable local state."""
    api = API()
    server = ThreadingHTTPServer(("127.0.0.1", 0), api.handler())
    thread = threading.Thread(target=server.serve_forever, daemon=True)
    thread.start()
    evidence: dict[str, Any] = {
        "transitions": [],
        "same_type_updates": [],
        "prevent_destroy": False,
        "create_before_destroy": False,
        "type_changing_puts": 0,
    }
    try:
        with tempfile.TemporaryDirectory(prefix="xcsh-lb-lifecycle-") as temporary:
            root = Path(temporary)
            env = {
                key: value
                for key, value in os.environ.items()
                if not key.startswith(("XCSH_", "TF_VAR_", "TF_LOG"))
            }
            env.update({"TF_IN_AUTOMATION": "1", "CHECKPOINT_DISABLE": "1"})
            if args.provider_bin:
                binary_dir = root / "provider"
                binary_dir.mkdir()
                binary = binary_dir / "terraform-provider-xcsh"
                binary.symlink_to(Path(args.provider_bin).resolve())
                rc = root / "terraform.rc"
                rc.write_text(
                    'provider_installation { dev_overrides { "f5-sales-demo/xcsh" = '
                    + json.dumps(str(binary_dir))
                    + " } direct {} }\n"
                )
                env["TF_CLI_CONFIG_FILE"] = str(rc)
            else:
                env.pop("TF_CLI_CONFIG_FILE", None)

            def run(
                *command: str, succeeds: bool = True
            ) -> subprocess.CompletedProcess[str]:
                result = subprocess.run(  # noqa: S603 -- fixed arguments without shell
                    [args.terraform, *command],
                    cwd=root,
                    env=env,
                    capture_output=True,
                    text=True,
                    check=False,
                    timeout=180,
                )
                if succeeds and result.returncode:
                    message = (
                        f"Terraform {' '.join(command)} failed:\n"
                        f"{result.stdout}\n{result.stderr}"
                    )
                    raise RuntimeError(message)
                return result

            def config(
                selection: str,
                setting: int = 80,
                lifecycle: str = "",
                certificate: str = "",
            ) -> None:
                block = ""
                if selection != "omitted":
                    cert = (
                        f'tls_cert_params {{\n certificates {{\n name = "{certificate}"\n'
                        'namespace = "default"\n }\n }'
                        if certificate
                        else ""
                    )
                    block = f"{selection} {{\n port = {setting}\n {cert}\n }}"
                version = (
                    f'version = "{args.registry_version}"'
                    if args.registry_version
                    else ""
                )
                (root / "main.tf").write_text(
                    f"""
terraform {{
  required_providers {{
    xcsh = {{ source = "f5-sales-demo/xcsh" {version} }}
  }}
}}
provider "xcsh" {{
  api_url = "http://127.0.0.1:{server.server_port}"
  api_token = "synthetic-loopback-token"
}}
resource "xcsh_http_loadbalancer" "test" {{
  name = "synthetic-lifecycle-lb"
  namespace = "default"
  domains = ["test.example.com"]
  {block}
  {lifecycle}
}}
"""
                )

            def plan() -> tuple[dict[str, Any], str]:
                output = run("plan", "-input=false", "-no-color", "-out=change.tfplan")
                data = json.loads(run("show", "-json", "change.tfplan").stdout)
                change = next(
                    item
                    for item in data["resource_changes"]
                    if item["address"] == "xcsh_http_loadbalancer.test"
                )
                return (change, output.stdout)

            def apply() -> None:
                run(
                    "apply",
                    "-input=false",
                    "-no-color",
                    "-auto-approve",
                    "change.tfplan",
                )

            def destroy() -> None:
                run("destroy", "-input=false", "-no-color", "-auto-approve")
                require(api.object is None, "lifecycle acceptance failed")

            config("http")
            if not args.provider_bin:
                run("init", "-input=false", "-no-color")
                evidence["registry_lock"] = (root / ".terraform.lock.hcl").read_text()
            types = ("http", "https", "https_auto_cert")
            for old in types:
                for new in types:
                    if old == new:
                        continue
                    config(old)
                    change, _ = plan()
                    require(change["change"]["actions"] == ["create"], change)
                    apply()
                    uid = api.current()["metadata"]["uid"]
                    start = len(api.writes)
                    config(new)
                    change, output = plan()
                    require(change["change"]["actions"] == ["delete", "create"], change)
                    paths = change["change"]["replace_paths"]
                    require([old] in paths and [new] in paths, paths)
                    require("must be replaced" in output, output)
                    require("Load Balancer Type Requires Replacement" in output, output)
                    apply()
                    require(
                        api.writes[start:] == ["DELETE", "POST"], api.writes[start:]
                    )
                    require(
                        api.current()["metadata"]["uid"] != uid,
                        "lifecycle acceptance failed",
                    )
                    evidence["transitions"].append(
                        {
                            "old": old,
                            "new": new,
                            "replace_paths": paths,
                            "writes": api.writes[start:],
                            "uid_changed": True,
                        }
                    )
                    destroy()
            for selected in types:
                config(selected, 80)
                plan()
                apply()
                uid = api.current()["metadata"]["uid"]
                start = len(api.writes)
                config(selected, 8080)
                change, _ = plan()
                require(change["change"]["actions"] == ["update"], change)
                apply()
                require(api.writes[start:] == ["PUT"], "lifecycle acceptance failed")
                require(
                    api.current()["metadata"]["uid"] == uid,
                    "lifecycle acceptance failed",
                )
                change, _ = plan()
                require(change["change"]["actions"] == ["no-op"], change)
                evidence["same_type_updates"].append(selected)
                destroy()
            config("https", certificate="synthetic-cert-one")
            plan()
            apply()
            uid = api.current()["metadata"]["uid"]
            start = len(api.writes)
            config("https", certificate="synthetic-cert-two")
            change, _ = plan()
            require(change["change"]["actions"] == ["update"], change)
            apply()
            require(api.writes[start:] == ["PUT"], api.writes[start:])
            require(
                api.current()["metadata"]["uid"] == uid,
                "certificate rotation changed UID",
            )
            evidence["certificate_rotation"] = True
            destroy()
            config("omitted")
            plan()
            apply()
            config("http")
            change, _ = plan()
            require(change["change"]["actions"] == ["delete", "create"], change)
            require(["http"] in change["change"]["replace_paths"], change)
            apply()
            config("omitted")
            change, _ = plan()
            require(change["change"]["actions"] == ["delete", "create"], change)
            apply()
            evidence["omitted_selection"] = True
            destroy()
            config("http")
            plan()
            apply()
            start = len(api.writes)
            config("https", lifecycle="lifecycle { prevent_destroy = true }")
            blocked_plan = run("plan", "-input=false", "-no-color", succeeds=False)
            require(
                blocked_plan.returncode != 0
                and "prevent_destroy" in blocked_plan.stdout + blocked_plan.stderr,
                "lifecycle acceptance failed",
            )
            require(not api.writes[start:], "lifecycle acceptance failed")
            evidence["prevent_destroy"] = True
            config("https", lifecycle="lifecycle { create_before_destroy = true }")
            change, _ = plan()
            require(change["change"]["actions"] == ["create", "delete"], change)
            require(not api.writes[start:], "lifecycle acceptance failed")
            evidence["create_before_destroy"] = True
            config("http")
            destroy()
            require(api.type_changing_puts == 0, "lifecycle acceptance failed")
            evidence["type_changing_puts"] = api.type_changing_puts
    finally:
        server.shutdown()
        server.server_close()
        thread.join()
    if args.evidence:
        Path(args.evidence).write_text(
            json.dumps(evidence, indent=2) + "\n", encoding="utf-8"
        )
    print(json.dumps(evidence, sort_keys=True))


def main() -> None:
    """Parse CLI options and run loopback acceptance."""
    parser = argparse.ArgumentParser(description=__doc__)
    source = parser.add_mutually_exclusive_group(required=True)
    source.add_argument("--provider-bin")
    source.add_argument("--registry-version")
    parser.add_argument("--terraform", default="terraform")
    parser.add_argument("--evidence")
    exercise(parser.parse_args())


if __name__ == "__main__":
    main()
