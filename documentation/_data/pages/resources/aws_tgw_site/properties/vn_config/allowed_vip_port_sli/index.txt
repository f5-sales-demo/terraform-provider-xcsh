---
page_title: "vn_config.allowed_vip_port_sli"
subcategory: ""
description: "vn_config.allowed_vip_port_sli for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 4583, "body_sha256": "sha256:d8167c528ff568e88cad52a7d13ad167f5555a2432c4b34ffed830859a5c1202", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "documentation/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["vn_config", "allowed_vip_port_sli"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.allowed_vip_port_sli for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port_sli

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- vn_config.allowed_vip_port_sli

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

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_https_port/): complete subsection reference.

## Next pages

- [vn_config.allowed_vip_port_sli.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/)
- [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/disable_allowed_vip_port/)
- [vn_config.allowed_vip_port_sli.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_https_port/)
- [vn_config.allowed_vip_port_sli.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_port/)
- [vn_config.allowed_vip_port_sli.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_https_port/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
