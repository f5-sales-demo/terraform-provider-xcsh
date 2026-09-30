---
page_title: "peers.external.interface_list"
subcategory: ""
description: "peers.external.interface_list for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1260, "body_sha256": "sha256:2f3785838bbfaf4bc695a38ba65431b03e1f1245ac06e3087a3d342263aaa461", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:interface_list:interfaces"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external", "path": "docs/guides/resources--bgp--properties--peers--external--interface_list.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.interface_list for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# peers.external.interface_list

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.external](resources--bgp--properties--peers--external.md)
- peers.external.interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interfaces](resources--bgp--properties--peers--external--interface_list--interfaces.md): complete subsection reference.

## Next pages

- [peers.external.interface_list.interfaces](resources--bgp--properties--peers--external--interface_list--interfaces.md)
- [peers.external](resources--bgp--properties--peers--external.md)
- [xcsh_bgp](../resources/bgp.md)
