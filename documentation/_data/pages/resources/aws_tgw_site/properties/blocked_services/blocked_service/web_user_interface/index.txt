---
page_title: "blocked_services.blocked_service.web_user_interface"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["blocked services blocked service web user interface"], "body_bytes": 1206, "body_sha256": "sha256:8476bf1c4600b861a2656fcd3be6a535a60ce73e1cba2f89d68d538841df3f29", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:web_user_interface", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service", "path": "documentation/resources/aws_tgw_site/properties/blocked_services/blocked_service/web_user_interface/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0331001001211212-0112333223132230-1202112232020332-1000310030213223-2121121313201202-1020113330330100-1201303212013222-3322211033220310", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_services", "blocked_service", "web_user_interface"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/blocked_services/blocked_service/web_user_interface/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services.blocked_service.web_user_interface

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/)
- [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/)
- blocked_services.blocked_service.web_user_interface

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
web_user_interface = {}
```

This is an empty object or choice marker. It has no direct properties.
