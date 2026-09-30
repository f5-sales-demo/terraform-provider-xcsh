---
page_title: "allow_all_response_codes"
subcategory: "Security"
description: "allow_all_response_codes for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1288, "body_sha256": "sha256:f6b4d37310f4c0c7892f1ae397d3c90a9f32e9a008a17cf9953ce517f938bad8", "canonical_id": "xcsh-docs:resources:app_firewall:properties:allow_all_response_codes", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:allow_all_response_codes", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "docs/guides/resources--app_firewall--properties--allow_all_response_codes.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_all_response_codes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/allow_all_response_codes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_all_response_codes for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allow_all_response_codes

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- allow_all_response_codes

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: allow\_all\_response\_codes, allowed\_response\_codes\] Configuration parameter for allow
all response codes. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

- [allow_all_response_codes](resources--app_firewall--properties--allow_all_response_codes.md#section)
- [allowed_response_codes](resources--app_firewall--properties--allowed_response_codes.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_response_codes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--app_firewall--reference.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
