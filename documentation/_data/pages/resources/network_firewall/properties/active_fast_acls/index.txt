---
page_title: "active_fast_acls"
subcategory: "Security"
description: "active_fast_acls for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1935, "body_sha256": "sha256:f8513aa398ec8bc069764d488b69da8738830a828b2a26c41995e4e2b87b7898", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls"], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "documentation/resources/network_firewall/properties/active_fast_acls/index.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["active_fast_acls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_fast_acls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_fast_acls for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# active_fast_acls

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- active_fast_acls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Upstream description:

List of Fast ACL(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fast_acls")}
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

OneOf alternatives in this subsection:

- [active_fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/#section)
- [disable_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/disable_fast_acl/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_fast_acls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/fast_acls/): complete subsection reference.

## Next pages

- [active_fast_acls.fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/fast_acls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
