---
page_title: "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations"
subcategory: "Load Balancing"
description: "Configuration parameter for code base integrations"
xcsh_docs: {"aliases": ["single lb app enable discovery api discovery from code scan code base integrations"], "body_bytes": 4377, "body_sha256": "sha256:6f8b8091eefe77dd8525e6b04d036e048d58a82c6cb3653ea4270c9e39d47476", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:code_base_integration", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations"], "schema_version": 1, "sections": [{"aliases": ["single lb app enable discovery api discovery from code scan code base integrations all repos"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations", "all_repos"], "syntax": "attribute", "type": "object"}, {"aliases": ["single lb app enable discovery api discovery from code scan code base integrations code base integration"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:code_base_integration", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration--name", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:code_base_integration", "type": "requires"}], "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations", "code_base_integration"], "syntax": "block", "type": "object"}, {"aliases": ["single lb app enable discovery api discovery from code scan code base integrations selected repos"], "anchor": "section", "description": "Select which API repositories represent the LB applications.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-single_lb_app--enable_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos:RequiredObjectAttributes:api_code_repo", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "type": "requires"}], "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan", "code_base_integrations", "selected_repos"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration parameter for code base integrations", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [all_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/all_repos/): complete subsection reference.

- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/code_base_integration/): complete subsection reference.

- [selected_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/all_repos/)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/code_base_integration/)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/code_base_integrations/selected_repos/)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
