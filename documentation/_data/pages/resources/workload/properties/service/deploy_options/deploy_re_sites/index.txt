---
page_title: "service.deploy_options.deploy_re_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge sites."
xcsh_docs: {"aliases": ["service deploy options deploy re sites"], "body_bytes": 1411, "body_sha256": "sha256:7084458fe2c3b43afe745e706f8534a84f2fb64e984a3bb7dd592db1e0e32a14", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_sites:site"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_sites", "parent_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "path": "documentation/resources/workload/properties/service/deploy_options/deploy_re_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0333112330330320-0311323321321100-3312230001320320-3323132011313121-3113201331330210-0310223030120322-3312212131002021-2032113220120030", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.deploy_options.deploy_re_sites:RequiredObjectAttributes:site", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_sites:site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "deploy_options", "deploy_re_sites"], "schema_version": 1, "sections": [{"aliases": ["service deploy options deploy re sites site"], "anchor": "section", "description": "Which regional edge sites should this workload be deployed.", "document_id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_sites:site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-service--deploy_options--deploy_re_sites--site--name", "enforcement": "provider-schema", "group": "service.deploy_options.deploy_re_sites.site:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_sites:site", "type": "requires"}], "schema_path": ["service", "deploy_options", "deploy_re_sites", "site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/deploy_re_sites/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines a way to deploy a workload on specific Regional Edge sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.deploy_re_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/)
- service.deploy_options.deploy_re_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to deploy a workload on specific Regional Edge sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/deploy_re_sites/site/): complete subsection reference.
