---
page_title: "allow_list"
subcategory: "Security"
description: "URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)"
xcsh_docs: {"aliases": ["allow list"], "body_bytes": 3284, "body_sha256": "sha256:0abd08c8b01e73d964c74a3ba32e3e611461dae18c60417ec532e228f22adda2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_allow", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_deny", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_next_policy", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:dest_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:tls_list"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:reference", "path": "documentation/data-sources/forward_proxy_policy/properties/allow_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allow_list"], "schema_version": 1, "sections": [{"aliases": ["default action allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_allow", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "default_action_allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["default action deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "default_action_deny"], "syntax": "attribute", "type": "object"}, {"aliases": ["default action next policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_next_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_list", "default_action_next_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["dest list"], "anchor": "section", "description": "L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:dest_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["allow_list", "dest_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["http list"], "anchor": "section", "description": "URLs for HTTP connections.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["allow_list", "http_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls list"], "anchor": "section", "description": "Domains in SNI for TLS connections.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:tls_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["allow_list", "tls_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/allow_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- allow_list

<a id="section"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

## Direct properties

- [default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/default_action_allow/): complete subsection reference.

- [default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/default_action_deny/): complete subsection reference.

- [default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/default_action_next_policy/): complete subsection reference.

- [dest_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/dest_list/): complete subsection reference.

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/http_list/): complete subsection reference.

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/tls_list/): complete subsection reference.

## Next pages

- [allow_list.default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/default_action_allow/)
- [allow_list.default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/default_action_deny/)
- [allow_list.default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/default_action_next_policy/)
- [allow_list.dest_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/dest_list/)
- [allow_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/http_list/)
- [allow_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/allow_list/tls_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
