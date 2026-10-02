---
page_title: "http_protocol_options.http_protocol_enable_v1_only"
subcategory: ""
description: "HTTP/1.1 Protocol OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["http protocol options http protocol enable v1 only"], "body_bytes": 1743, "body_sha256": "sha256:3214825fd887d031f06645bc0b63aaafe2ebdc381c4ffd8732303002c9dc16a6", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "parent_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options", "path": "documentation/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3210022102331030-1102301122203233-1032122010103223-3301302200330220-3133302113323110-0303211131002113-0323210330020030-2332000211221030", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only"], "schema_version": 1, "sections": [{"aliases": ["header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,preserve_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:default_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options.http_protocol_enable_v1_only.header_transformation:ConflictingObjectAttributes:preserve_case_header_transformation,proper_case_header_transformation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation", "type": "conflicts"}], "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options.http_protocol_enable_v1_only

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/)
- http_protocol_options.http_protocol_enable_v1_only

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

## Direct properties

- [header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/): complete subsection reference.

## Next pages

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
