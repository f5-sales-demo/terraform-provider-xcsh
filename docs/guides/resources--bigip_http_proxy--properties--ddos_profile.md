---
page_title: "ddos_profile"
subcategory: ""
description: "ddos_profile for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1612, "body_sha256": "sha256:caddf37094647afc4d1642142167b8d41dae9b7327a35fe387218012cb3272ad", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:disable_ddos_mitigation", "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile:enable_ddos_mitigation"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "docs/guides/resources--bigip_http_proxy--properties--ddos_profile.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ddos_profile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/ddos_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ddos_profile for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ddos_profile

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- ddos_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Upstream description:

BIG-IP DDoS Protection Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
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
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_ddos_mitigation](resources--bigip_http_proxy--properties--ddos_profile--disable_ddos_mitigation.md): complete subsection reference.

- [enable_ddos_mitigation](resources--bigip_http_proxy--properties--ddos_profile--enable_ddos_mitigation.md): complete subsection reference.

## Next pages

- [ddos_profile.disable_ddos_mitigation](resources--bigip_http_proxy--properties--ddos_profile--disable_ddos_mitigation.md)
- [ddos_profile.enable_ddos_mitigation](resources--bigip_http_proxy--properties--ddos_profile--enable_ddos_mitigation.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
