---
page_title: "active_forward_proxy_policies"
subcategory: ""
description: "Ordered List of Forward Proxy Policies active."
xcsh_docs: {"aliases": ["active forward proxy policies"], "body_bytes": 1524, "body_sha256": "sha256:3067007651a567dc586411ee6d9b9a5bce328ac808ce25e9e4a5e60d54b0bac5", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:active_forward_proxy_policies:forward_proxy_policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/active_forward_proxy_policies/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0003323102113131-0103011023111231-1113001332012010-0301232233010331-0130302132021322-2321100013200013-0033323321132110-1121322121220021", "registry_path": "docs/guides/resources--proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["active_forward_proxy_policies"], "schema_version": 1, "sections": [{"aliases": ["active forward proxy policies forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:resources:proxy:properties:active_forward_proxy_policies:forward_proxy_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["active_forward_proxy_policies", "forward_proxy_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Ordered List of Forward Proxy Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["proxyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_forward_proxy_policies

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- active_forward_proxy_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, no\_forward\_proxy\_policy; Default:
no\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

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

OneOf alternatives in this subsection:

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/active_forward_proxy_policies/#section)
- [no_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/no_forward_proxy_policy/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/active_forward_proxy_policies/forward_proxy_policies/): complete subsection reference.
