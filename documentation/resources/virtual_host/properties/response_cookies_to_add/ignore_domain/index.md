---
page_title: "response_cookies_to_add.ignore_domain"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["response cookies to add ignore domain"], "body_bytes": 1049, "body_sha256": "sha256:7f2166828532920bc52dfcc9a238dc5a9e597df8d679465de4c0349c9de8ea68", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_domain", "parent_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "path": "documentation/resources/virtual_host/properties/response_cookies_to_add/ignore_domain/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3201100313033331-3021302030332320-1031022021211101-0112002311330312-1102000110001113-2323002121322130-3301322102003301-1102202322300212", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cookies_to_add", "ignore_domain"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/response_cookies_to_add/ignore_domain/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cookies_to_add.ignore_domain

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/)
- response_cookies_to_add.ignore_domain

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_domain = {}
```

This is an empty object or choice marker. It has no direct properties.
