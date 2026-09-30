---
page_title: "ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6"
subcategory: ""
description: "ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 2620, "body_sha256": "sha256:aab517d72bae6e65b6f9472b5c0e191808ea01219a7db66bfdc21aad9c553183", "canonical_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack:ipv6", "child_ids": [], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack:ipv6", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ike_parameters:rm_ip_address:dual_stack", "path": "docs/guides/data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters", "rm_ip_address", "dual_stack", "ipv6"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv6/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6 for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md)
- [Property reference](data-sources--external_connector--reference.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
- [ipsec.ike_parameters](data-sources--external_connector--properties--ipsec--ike_parameters.md)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack.md)
- ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6

<a id="section"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="schema-ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

## Next pages

- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack.md)
- [xcsh_external_connector](../data-sources/external_connector.md)
