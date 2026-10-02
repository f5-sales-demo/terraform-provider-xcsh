---
page_title: "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations"
subcategory: "Load Balancing"
description: "Configuration parameter for code base integrations"
xcsh_docs: {"aliases": ["enable api discovery api discovery from code scan code base integrations"], "body_bytes": 4064, "body_sha256": "sha256:088b9f9d9b17778309768ee00996c0538c54d6b5a22d8150e1d873953cbbe63e", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:code_base_integration", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "path": "documentation/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations"], "schema_version": 1, "sections": [{"aliases": ["all repos"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations", "all_repos"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:code_base_integration", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration--name", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:code_base_integration", "type": "requires"}], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations", "code_base_integration"], "syntax": "block", "type": "object"}, {"aliases": ["selected repos"], "anchor": "section", "description": "Select which API repositories represent the LB applications.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos:RequiredObjectAttributes:api_code_repo", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "type": "requires"}], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for code base integrations", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/)
- [enable_api_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for code base integrations.

Upstream description:

Configuration parameter for code base integrations

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_repos",
    "selected_repos")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/all_repos/): complete subsection reference.

- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/code_base_integration/): complete subsection reference.

- [selected_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/all_repos/)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/code_base_integration/)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/)
- [enable_api_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
