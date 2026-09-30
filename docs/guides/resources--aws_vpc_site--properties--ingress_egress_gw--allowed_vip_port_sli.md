---
page_title: "ingress_egress_gw.allowed_vip_port_sli"
subcategory: "Infrastructure"
description: "ingress_egress_gw.allowed_vip_port_sli for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3905, "body_sha256": "sha256:d9ee2f3a8276195bf2265eba1f07c34dc4ff1f8772bcb00c20aa3d26bd6b1d94", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:custom_ports", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:disable_allowed_vip_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_http_https_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_http_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli:use_https_port"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "allowed_vip_port_sli"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.allowed_vip_port_sli for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.allowed_vip_port_sli

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.allowed_vip_port_sli

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port_sli {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_ports](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--custom_ports.md): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--disable_allowed_vip_port.md): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_http_https_port.md): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_http_port.md): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_https_port.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.allowed_vip_port_sli.custom_ports](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--custom_ports.md)
- [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--disable_allowed_vip_port.md)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_http_https_port.md)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_http_port.md)
- [ingress_egress_gw.allowed_vip_port_sli.use_https_port](resources--aws_vpc_site--properties--ingress_egress_gw--allowed_vip_port_sli--use_https_port.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
