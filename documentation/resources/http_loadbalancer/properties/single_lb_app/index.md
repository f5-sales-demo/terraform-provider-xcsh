---
page_title: "single_lb_app"
subcategory: "Load Balancing"
description: "Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs."
xcsh_docs: {"aliases": ["single lb app"], "body_bytes": 3127, "body_sha256": "sha256:b672a339d04f27fbfc87e5bd5894319a7aef8ad8c52dd6f599a79f5c0df80a6b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_discovery", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_malicious_user_detection", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_malicious_user_detection"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-026.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app:ConflictingObjectAttributes:disable_discovery,enable_discovery", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_discovery", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app:ConflictingObjectAttributes:disable_malicious_user_detection,enable_malicious_user_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_malicious_user_detection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app:ConflictingObjectAttributes:disable_discovery,enable_discovery", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app:ConflictingObjectAttributes:disable_malicious_user_detection,enable_malicious_user_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_malicious_user_detection", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app"], "schema_version": 1, "sections": [{"aliases": ["disable discovery"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_discovery", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "disable_discovery"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable malicious user detection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_malicious_user_detection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "disable_malicious_user_detection"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable discovery"], "anchor": "section", "description": "Specifies the settings used for API discovery.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery:ConflictingObjectAttributes:custom_api_auth_discovery,default_api_auth_discovery", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery:ConflictingObjectAttributes:custom_api_auth_discovery,default_api_auth_discovery", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:default_api_auth_discovery", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery:ConflictingObjectAttributes:disable_learn_from_redirect_traffic,enable_learn_from_redirect_traffic", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:disable_learn_from_redirect_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "single_lb_app.enable_discovery:ConflictingObjectAttributes:disable_learn_from_redirect_traffic,enable_learn_from_redirect_traffic", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:enable_learn_from_redirect_traffic", "type": "conflicts"}], "schema_path": ["single_lb_app", "enable_discovery"], "syntax": "block", "type": "object"}, {"aliases": ["enable malicious user detection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_malicious_user_detection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_malicious_user_detection"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- single_lb_app

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_discovery",
    "enable_discovery"),
  validators.ConflictingObjectAttributes("disable_malicious_user_detection",
    "enable_malicious_user_detection")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_choice": "[\"disable_discovery\",\"enable_discovery\"]",
  "x-ves-oneof-field-malicious_user_detection_choice": "[\"disable_malicious_user_detection\",\"enable_malicious_user_detection\"]"
}
```

Terraform syntax:

```terraform
single_lb_app {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/disable_discovery/): complete subsection reference.

- [disable_malicious_user_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/disable_malicious_user_detection/): complete subsection reference.

- [enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/): complete subsection reference.

- [enable_malicious_user_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_malicious_user_detection/): complete subsection reference.

## Next pages

- [single_lb_app.disable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/disable_discovery/)
- [single_lb_app.disable_malicious_user_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/disable_malicious_user_detection/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_malicious_user_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_malicious_user_detection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
