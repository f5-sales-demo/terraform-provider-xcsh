---
page_title: "response_cookies_to_add.ignore_domain"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["response cookies to add ignore domain"], "body_bytes": 1002, "body_sha256": "sha256:eafa84dcf1fc3c92fed8b26242800ec6bc6313088dbb69904519d1d8a000ad45", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:response_cookies_to_add:ignore_domain", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:response_cookies_to_add", "path": "documentation/data-sources/virtual_host/properties/response_cookies_to_add/ignore_domain/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3130310133231002-1131000313113312-2030022200200233-2011302002303013-3303011201121220-3202203133303302-2011203131131013-2332211220103303", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cookies_to_add", "ignore_domain"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/response_cookies_to_add/ignore_domain/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cookies_to_add.ignore_domain

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/response_cookies_to_add/)
- response_cookies_to_add.ignore_domain

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore domain.

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

This is an empty object or choice marker. It has no direct properties.
