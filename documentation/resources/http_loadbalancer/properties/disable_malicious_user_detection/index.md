---
page_title: "disable_malicious_user_detection"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable malicious user detection"], "body_bytes": 1867, "body_sha256": "sha256:2beeecee5e7c3b9ab6b9f1b02e12636b7b2af1b18494ea3ed21a81d162fcc8f2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:disable_malicious_user_detection", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/disable_malicious_user_detection/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3012231223303232-2102332203003323-3102212131101310-2012202210031310-2221201202003333-1132132220302101-0321000202112303-1321313331130122", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_malicious_user_detection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/disable_malicious_user_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_malicious_user_detection

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- disable_malicious_user_detection

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.
Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/disable_malicious_user_detection/#section)
- [enable_malicious_user_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_malicious_user_detection/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
