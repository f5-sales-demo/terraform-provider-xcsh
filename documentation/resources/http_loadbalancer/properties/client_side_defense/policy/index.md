---
page_title: "client_side_defense.policy"
subcategory: "Load Balancing"
description: "client_side_defense.policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3756, "body_sha256": "sha256:4368945d8799fdfb3139fcab78852bd6188acb186a7f2c593b24bf54e81def3c", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense", "path": "documentation/resources/http_loadbalancer/properties/client_side_defense/policy/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["client_side_defense", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/)
- client_side_defense.policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/disable_js_insert/): complete subsection reference.

- [js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages/): complete subsection reference.

- [js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/): complete subsection reference.

## Next pages

- [client_side_defense.policy.disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/disable_js_insert/)
- [client_side_defense.policy.js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages/)
- [client_side_defense.policy.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
