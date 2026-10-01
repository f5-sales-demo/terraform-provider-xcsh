---
page_title: "single_lb_app"
subcategory: "Load Balancing"
description: "single_lb_app for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3127, "body_sha256": "sha256:b672a339d04f27fbfc87e5bd5894319a7aef8ad8c52dd6f599a79f5c0df80a6b", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_discovery", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_malicious_user_detection", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_malicious_user_detection"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["single_lb_app"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
