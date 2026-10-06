---
page_title: "single_lb_app.enable_discovery.api_discovery_from_code_scan"
subcategory: "Load Balancing"
description: "Select Code Base and Repositories."
xcsh_docs: {"aliases": ["single lb app enable discovery api discovery from code scan"], "body_bytes": 1615, "body_sha256": "sha256:c54bda19b9a7a11ad962e34d6651e40859cbd07808bc0a6a0de6b953251887e2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_discovery_from_code_scan:RequiredObjectAttributes:code_base_integrations", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan"], "schema_version": 1, "sections": [{"aliases": ["single lb app enable discovery api discovery from code scan code base integrations"], "anchor": "section", "description": "Configuration parameter for code base integrations", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "type": "conflicts"}], "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Select Code Base and Repositories.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_discovery_from_code_scan

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- single_lb_app.enable_discovery.api_discovery_from_code_scan

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Code Base and Repositories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

## Direct properties

- [code_base_integrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/): complete subsection reference.
