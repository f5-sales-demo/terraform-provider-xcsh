---
page_title: "ipsec.ike_parameters.rm_ip_address.ipv4"
subcategory: ""
description: "ipsec.ike_parameters.rm_ip_address.ipv4 for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1904, "body_sha256": "sha256:3dc48375f57e5eb5be0e0c1a1ad0c74a473c1a79a51fae1c41c772cfebb88f32", "canonical_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:ipv4", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address", "path": "docs/guides/data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv4.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters", "rm_ip_address", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters.rm_ip_address.ipv4 for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ipsec.ike_parameters.rm_ip_address.ipv4

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md)
- [Property reference](data-sources--external_connector--reference.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
- [ipsec.ike_parameters](data-sources--external_connector--properties--ipsec--ike_parameters.md)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md)
- ipsec.ike_parameters.rm_ip_address.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

## Direct properties

<a id="schema-ipsec--ike_parameters--rm_ip_address--ipv4--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md)
- [xcsh_external_connector](../data-sources/external_connector.md)
