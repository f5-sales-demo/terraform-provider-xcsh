---
page_title: "email"
subcategory: ""
description: "email for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2392, "body_sha256": "sha256:7cc25808e9c3e34048b51a1038c706c2f7f18e6ec2b8f765629c9a3817f4a40a", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:email", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:email", "parent_id": "xcsh-docs:resources:alert_receiver:reference", "path": "docs/guides/resources--alert_receiver--properties--email.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["email"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/email/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "email for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# email

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- email

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: email, opsgenie, pagerduty, slack, sms, webhook\] Email Configuration.

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

- [email](resources--alert_receiver--properties--email.md#section)
- [opsgenie](resources--alert_receiver--properties--opsgenie.md#section)
- [pagerduty](resources--alert_receiver--properties--pagerduty.md#section)
- [slack](resources--alert_receiver--properties--slack.md#section)
- [sms](resources--alert_receiver--properties--sms.md#section)
- [webhook](resources--alert_receiver--properties--webhook.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
email {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-email--email"></a>

### email property

Type: `"string"`. Optional.

Email. Email ID of the user.

Upstream description:

Email ID of the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

## Next pages

- [Property reference](resources--alert_receiver--reference.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
