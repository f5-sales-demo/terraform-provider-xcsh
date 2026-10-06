---
page_title: "stateful_service.advertise_options.advertise_custom.advertise_where"
subcategory: "Container"
description: "Where should this load balancer be available."
xcsh_docs: {"aliases": ["stateful service advertise options advertise custom advertise where"], "body_bytes": 3145, "body_sha256": "sha256:4b7ca9dac0b8cf1d2ecc773a6e053efa59ab8a1732da800694cc0c991845b035", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:site", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:virtual_site", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_custom/advertise_where/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310", "registry_path": "docs/guides/resources--workload--reference--group-018.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "advertise_where"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise custom advertise where site"], "anchor": "section", "description": "This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "advertise_where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise custom advertise where virtual site"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "advertise_where", "virtual_site"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise custom advertise where vk8s service"], "anchor": "section", "description": "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "conflicts"}], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "advertise_where", "vk8s_service"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Where should this load balancer be available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/advertise_where/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/advertise_where/virtual_site/): complete subsection reference.

- [vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/advertise_where/vk8s_service/): complete subsection reference.
