---
page_title: "service.advertise_options.advertise_custom.advertise_where"
subcategory: "Container"
description: "Where should this load balancer be available."
xcsh_docs: {"aliases": ["service advertise options advertise custom advertise where"], "body_bytes": 3076, "body_sha256": "sha256:1d85d963c45818b1aab0cf88873f92a44a7415432b73f8b207bcceb0a51aa3d4", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:site", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:virtual_site", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise custom advertise where site"], "anchor": "section", "description": "This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise custom advertise where virtual site"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where", "virtual_site"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise custom advertise where vk8s service"], "anchor": "section", "description": "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where", "vk8s_service"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Where should this load balancer be available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- service.advertise_options.advertise_custom.advertise_where

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/virtual_site/): complete subsection reference.

- [vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/): complete subsection reference.
