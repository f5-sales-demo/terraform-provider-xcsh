---
page_title: "peers.external.family_inet.enable"
subcategory: ""
description: "peers.external.family_inet.enable for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1315, "body_sha256": "sha256:55aaf931d28e1faa51f82bbd6545676e56690364c6b78c912a75e32869c28632", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet", "path": "docs/guides/resources--bgp--properties--peers--external--family_inet--enable.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/family_inet/enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external.family_inet.enable for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.external](resources--bgp--properties--peers--external.md)
- [peers.external.family_inet](resources--bgp--properties--peers--external--family_inet.md)
- peers.external.family_inet.enable

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

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
enable {
  # Configure direct properties listed below.
}
```

## Direct properties

- [aggregation](resources--bgp--properties--peers--external--family_inet--enable--aggregation.md): complete subsection reference.

## Next pages

- [peers.external.family_inet.enable.aggregation](resources--bgp--properties--peers--external--family_inet--enable--aggregation.md)
- [peers.external.family_inet](resources--bgp--properties--peers--external--family_inet.md)
- [xcsh_bgp](../resources/bgp.md)
