---
page_title: "tgw_security"
subcategory: ""
description: "tgw_security for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 5797, "body_sha256": "sha256:f69c9990219c819a7e954e30336c1ff5954a4da8623f9b76901d7f37a5d5ebb1", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_network_policies", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_network_policy"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/tgw_security/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["tgw_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tgw_security for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
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

- [active_east_west_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/): complete subsection reference.

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_network_policies/): complete subsection reference.

- [east_west_service_policy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/east_west_service_policy_allow_all/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/forward_proxy_allow_all/): complete subsection reference.

- [no_east_west_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_east_west_policy/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_forward_proxy/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_network_policy/): complete subsection reference.

## Next pages

- [tgw_security.active_east_west_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/)
- [tgw_security.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/)
- [tgw_security.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/)
- [tgw_security.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_network_policies/)
- [tgw_security.east_west_service_policy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/east_west_service_policy_allow_all/)
- [tgw_security.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/forward_proxy_allow_all/)
- [tgw_security.no_east_west_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_east_west_policy/)
- [tgw_security.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_forward_proxy/)
- [tgw_security.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
