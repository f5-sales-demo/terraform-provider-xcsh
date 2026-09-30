---
page_title: "peers.external.family_inet"
subcategory: ""
description: "peers.external.family_inet for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 2049, "body_sha256": "sha256:42cd1a3144a91c6b16edaaacbc80a8843c2e8ad356918a85bbe005bc8f08ba05", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:family_inet:disable_spec", "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external", "path": "documentation/resources/bgp/properties/peers/external/family_inet/index.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["peers", "external", "family_inet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/family_inet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.family_inet for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# peers.external.family_inet

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- peers.external.family_inet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
family_inet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/): complete subsection reference.

## Next pages

- [peers.external.family_inet.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/disable_spec/)
- [peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
