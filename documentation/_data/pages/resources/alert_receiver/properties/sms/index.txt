---
page_title: "sms"
subcategory: ""
description: "SMS Configuration."
xcsh_docs: {"aliases": ["sms"], "body_bytes": 1904, "body_sha256": "sha256:aad8313d77a2600ae3bd7539aa14f31c5963abd40c24d590fed4c05ed2d829f1", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:sms", "parent_id": "xcsh-docs:resources:alert_receiver:reference", "path": "documentation/resources/alert_receiver/properties/sms/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2321322330122322-3323101322120202-0332000331221112-1333103013101001-3330200333023112-3123201202122121-1121033313113332-2312331010332202", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sms"], "schema_version": 1, "sections": [{"aliases": ["sms contact number"], "anchor": "schema-sms--contact_number", "description": "Contact number of the user in ITU E.164 format", "document_id": "xcsh-docs:resources:alert_receiver:properties:sms", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sms", "contact_number"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/sms/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SMS Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sms

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- sms

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SMS Configuration.

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
sms {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-sms--contact_number"></a>

### contact_number property

Type: `"string"`. Optional.

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\].

Upstream description:

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
