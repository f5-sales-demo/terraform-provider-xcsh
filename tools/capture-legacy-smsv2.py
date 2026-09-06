#!/usr/bin/env python3
"""Capture pinned legacy SMSv2 requests using an isolated loopback API.

Requires Terraform, OpenSSL and the installed Linux amd64 Volterra 0.12.2 binary.
No cloud credentials or cloud resources are used. Terraform state and the
short-lived test certificate stay in a temporary directory and are removed.
"""

from __future__ import annotations

import argparse
import hashlib
import json
import logging
import os
import shutil
import ssl
import subprocess
import tempfile
import threading
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path
from typing import Any

LOGGER = logging.getLogger(__name__)

LEGACY_SHA256 = "821231f24ac1d4d95075493afc1a681d601888e5bff09859eaa71fb037c478e2"
SOURCE_COMMIT = "22f029dbf14412c99502fe1daba829f7c3261017"
PLATFORMS = (
    "aws",
    "azure",
    "baremetal",
    "equinix",
    "gcp",
    "kvm",
    "nutanix",
    "oci",
    "openshift_virtualization",
    "openstack",
    "rseries",
    "vmware",
)


def capture(binary: Path, platform: str, scenario: str = "discovery") -> dict[str, Any]:
    """Execute the actual installed provider; retain only synthetic requests."""
    if hashlib.sha256(binary.read_bytes()).hexdigest() != LEGACY_SHA256:
        message = "legacy binary digest does not match pinned 0.12.2"
        raise ValueError(message)
    terraform = shutil.which("terraform")
    openssl = shutil.which("openssl")
    if not terraform or not openssl:
        message = "Terraform and OpenSSL must be installed"
        raise ValueError(message)
    requests: list[dict[str, Any]] = []
    objects: dict[str, Any] = {}

    class Handler(BaseHTTPRequestHandler):
        def log_message(self, *_args: Any) -> None:
            pass

        def respond(self, body: dict[str, Any]) -> None:
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps(body).encode())

        def do_POST(self) -> None:
            body = json.loads(
                self.rfile.read(int(self.headers.get("Content-Length", "0")))
            )
            requests.append({"method": "POST", "path": self.path, "body": body})
            name = body["metadata"]["name"]
            objects[name] = {
                **body,
                "system_metadata": {
                    "uid": "11111111-2222-4333-8444-555555555555",
                    "creation_timestamp": "2026-01-01T00:00:00Z",
                },
                "resource_version": "1",
            }
            self.respond(objects[name])

        def do_GET(self) -> None:
            name = self.path.rsplit("/", 1)[-1].split("?")[0]
            self.respond(objects.get(name, {"items": []}))

    with tempfile.TemporaryDirectory(prefix="legacy-smsv2-wire-") as directory:
        root = Path(directory)
        plugin_dir = root / "plugins"
        plugin_dir.mkdir()
        shutil.copy2(binary, plugin_dir / "terraform-provider-volterra_v0.12.2")
        certificate, key = root / "client.crt", root / "client.key"
        subprocess.run(  # noqa: S603 - resolved OpenSSL executable and fixed arguments
            [
                openssl,
                "req",
                "-x509",
                "-newkey",
                "rsa:2048",
                "-nodes",
                "-keyout",
                str(key),
                "-out",
                str(certificate),
                "-days",
                "1",
                "-subj",
                "/CN=localhost",
                "-addext",
                "subjectAltName=IP:127.0.0.1",
            ],
            check=True,
            capture_output=True,
            timeout=30,
        )
        key.chmod(0o600)
        server = ThreadingHTTPServer(("127.0.0.1", 0), Handler)
        tls = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
        tls.load_cert_chain(certificate, key)
        server.socket = tls.wrap_socket(server.socket, server_side=True)
        worker = threading.Thread(target=server.serve_forever, daemon=True)
        worker.start()
        cli_config = root / "terraform.rc"
        cli_config.write_text(
            'provider_installation { dev_overrides { "volterraedge/volterra" = '
            + json.dumps(str(plugin_dir))
            + " } direct {} }"
        )
        config = {
            "terraform": {
                "required_providers": {
                    "volterra": {
                        "source": "volterraedge/volterra",
                        "version": "0.12.2",
                    },
                }
            },
            "provider": {
                "volterra": {
                    "url": f"https://127.0.0.1:{server.server_port}/api",
                    "api_cert": str(certificate),
                    "api_key": str(key),
                    "api_ca_cert": str(certificate),
                }
            },
            "resource": {
                "volterra_securemesh_site_v2": {
                    "fixture": {
                        "name": "smsv2-parity-fixture",
                        "namespace": "system",
                        "description": "independent legacy wire fixture",
                        "block_all_services": True,
                        "logs_streaming_disabled": True,
                        "enable_ha": False,
                        "re_select": [{"geo_proximity": True}],
                        platform: [{"not_managed": [{}]}],
                        "software_settings": [
                            {
                                "os": [{"default_os_version": True}],
                                "sw": [{"default_sw_version": True}],
                            }
                        ],
                    }
                }
            },
        }
        if scenario != "discovery":
            interface = {
                "name": "parity-interface",
                "ethernet_interface": [{"device": "ens6", "mac": "02:00:00:00:00:06"}],
                "no_ipv6_address": True,
                "network_option": [{"site_local_inside_network": True}],
            }
            if scenario == "static-dns":
                interface["static_ip"] = [
                    {
                        "ip_address": "192.0.2.10/24",
                        "default_gw": "192.0.2.1",
                        "dns_server": "192.0.2.53",
                    }
                ]
            elif scenario == "dhcp-server":
                interface["dhcp_server"] = [
                    {
                        "automatic_from_start": True,
                        "dhcp_option82_tag": "parity-option82",
                        "fixed_ip_map": {"02:00:00:00:00:20": "192.0.2.20"},
                        "dhcp_networks": [
                            {
                                "network_prefix": "192.0.2.0/24",
                                "dgw_address": "192.0.2.1",
                                "dns_address": "192.0.2.53",
                                "pool_settings": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
                                "pools": [
                                    {
                                        "start_ip": "192.0.2.20",
                                        "end_ip": "192.0.2.30",
                                        "exclude": True,
                                    }
                                ],
                            }
                        ],
                    }
                ]
            elif scenario == "dhcpv6-server":
                interface.pop("no_ipv6_address")
                interface["dhcp_client"] = True
                interface["ipv6_auto_config"] = [
                    {
                        "router": [
                            {
                                "stateful": [
                                    {
                                        "dhcp_networks": [
                                            {
                                                "network_prefix": "2001:db8:1::/64",
                                                "pool_settings": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
                                                "pools": [
                                                    {
                                                        "start_ip": "2001:db8:1::20",
                                                        "end_ip": "2001:db8:1::30",
                                                        "exclude": True,
                                                    }
                                                ],
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ]
            else:
                message = f"unknown scenario: {scenario}"
                raise ValueError(message)
            config["resource"]["volterra_securemesh_site_v2"]["fixture"][platform] = [
                {
                    "not_managed": [
                        {
                            "node_list": [
                                {
                                    "hostname": "parity-node",
                                    "interface_list": [interface],
                                }
                            ]
                        }
                    ],
                }
            ]
        (root / "main.tf.json").write_text(json.dumps(config))
        env = {
            k: v
            for k, v in os.environ.items()
            if not k.startswith(("VES_", "VOLTERRA_", "XCSH_", "TF_"))
        }
        env["TF_CLI_CONFIG_FILE"] = str(cli_config)
        env["NO_PROXY"] = "127.0.0.1,localhost"
        env["no_proxy"] = env["NO_PROXY"]
        try:
            result = subprocess.run(  # noqa: S603 - resolved Terraform, isolated fixed config
                [terraform, "apply", "-auto-approve", "-input=false", "-no-color"],
                cwd=root,
                env=env,
                capture_output=True,
                text=True,
                timeout=90,
                check=False,
            )
            if result.returncode:
                message = f"loopback Terraform failed: {result.stdout}\n{result.stderr}"
                raise RuntimeError(message)
        finally:
            server.shutdown()
            server.server_close()
            worker.join(timeout=5)
    if len(requests) != 1:
        message = f"expected one create request, captured {len(requests)}"
        raise ValueError(message)
    if scenario == "discovery" and requests[0]["body"]["spec"].get(platform) != {
        "not_managed": {}
    }:
        message = "legacy platform discovery did not serialize as expected"
        raise ValueError(message)
    result = {
        "legacy_version": "0.12.2",
        "source_commit": SOURCE_COMMIT,
        "linux_amd64_binary_sha256": LEGACY_SHA256,
        "case": f"{platform}-{scenario}-software-defaults",
        "request": requests[0],
        "notes": [
            "Captured from the installed legacy binary against a loopback HTTPS server.",
            "enable_ha=false is omitted by the legacy SDK GetOk mapping.",
            "This is serialization evidence, not platform live acceptance.",
        ],
    }
    if scenario != "discovery":
        result["configuration"] = config["resource"]["volterra_securemesh_site_v2"][
            "fixture"
        ]
    return result


def main() -> None:
    """Capture one independently sourced discovery fixture."""
    logging.basicConfig(level=logging.INFO, format="%(message)s")
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--binary", type=Path, required=True)
    parser.add_argument("--platform", choices=PLATFORMS, required=True)
    parser.add_argument(
        "--scenario",
        choices=("discovery", "static-dns", "dhcp-server", "dhcpv6-server"),
        default="discovery",
    )
    parser.add_argument("--output", type=Path, required=True)
    args = parser.parse_args()
    result = capture(args.binary.resolve(strict=True), args.platform, args.scenario)
    args.output.write_text(json.dumps(result, indent=2) + "\n")
    LOGGER.info(
        "Captured %s legacy %s request from pinned binary", args.platform, args.scenario
    )


if __name__ == "__main__":
    main()
