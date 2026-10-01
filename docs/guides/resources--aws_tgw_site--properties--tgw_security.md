---
page_title: "tgw_security"
subcategory: ""
description: "tgw_security for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 4689, "body_sha256": "sha256:5a5930679ef6ef2c91e4d4f0658d2f6ea0ba5ffbaf1a7689c0c9a9c6c5e6edd7", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_network_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_network_policy"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "docs/guides/resources--aws_tgw_site--properties--tgw_security.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tgw_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- tgw_security

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Security Configuration for transit gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "east_west_service_policy_allow_all"),
  validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "no_east_west_policy"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("east_west_service_policy_allow_all",
    "no_east_west_policy"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy")}
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
  "x-ves-oneof-field-east_west_service_policy_choice": "[\"active_east_west_service_policies\",\"east_west_service_policy_allow_all\",\"no_east_west_policy\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]"
}
```

Terraform syntax:

```terraform
tgw_security {
  # Configure direct properties listed below.
}
```

## Direct properties

- [active_east_west_service_policies](resources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies.md): complete subsection reference.

- [active_enhanced_firewall_policies](resources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies.md): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies.md): complete subsection reference.

- [active_network_policies](resources--aws_tgw_site--properties--tgw_security--active_network_policies.md): complete subsection reference.

- [east_west_service_policy_allow_all](resources--aws_tgw_site--properties--tgw_security--east_west_service_policy_allow_all.md): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_tgw_site--properties--tgw_security--forward_proxy_allow_all.md): complete subsection reference.

- [no_east_west_policy](resources--aws_tgw_site--properties--tgw_security--no_east_west_policy.md): complete subsection reference.

- [no_forward_proxy](resources--aws_tgw_site--properties--tgw_security--no_forward_proxy.md): complete subsection reference.

- [no_network_policy](resources--aws_tgw_site--properties--tgw_security--no_network_policy.md): complete subsection reference.

## Next pages

- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--properties--tgw_security--active_east_west_service_policies.md)
- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--properties--tgw_security--active_enhanced_firewall_policies.md)
- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--properties--tgw_security--active_forward_proxy_policies.md)
- [tgw_security.active_network_policies](resources--aws_tgw_site--properties--tgw_security--active_network_policies.md)
- [tgw_security.east_west_service_policy_allow_all](resources--aws_tgw_site--properties--tgw_security--east_west_service_policy_allow_all.md)
- [tgw_security.forward_proxy_allow_all](resources--aws_tgw_site--properties--tgw_security--forward_proxy_allow_all.md)
- [tgw_security.no_east_west_policy](resources--aws_tgw_site--properties--tgw_security--no_east_west_policy.md)
- [tgw_security.no_forward_proxy](resources--aws_tgw_site--properties--tgw_security--no_forward_proxy.md)
- [tgw_security.no_network_policy](resources--aws_tgw_site--properties--tgw_security--no_network_policy.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
