---
page_title: "gre.gre_parameters.segment"
subcategory: ""
description: "Reference to Segment Object."
xcsh_docs: {"aliases": ["gre gre parameters segment"], "body_bytes": 1881, "body_sha256": "sha256:e94d3c50f76ad7507e9e235b77ac35d58f4a388c8f9f1612dd1474a06607f770", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment:refs"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "parent_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "path": "documentation/resources/external_connector/properties/gre/gre_parameters/segment/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3033011103100101-0331031313031322-1321223022000220-2121023221103112-1031123102133113-0310331122002210-2232000331012220-1200323030030230", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters.segment:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment:refs", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gre", "gre_parameters", "segment"], "schema_version": 1, "sections": [{"aliases": ["refs"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment:refs", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["gre", "gre_parameters", "segment", "refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Reference to Segment Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters.segment

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/)
- [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/)
- gre.gre_parameters.segment

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

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
segment {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/): complete subsection reference.

## Next pages

- [gre.gre_parameters.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/)
- [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
