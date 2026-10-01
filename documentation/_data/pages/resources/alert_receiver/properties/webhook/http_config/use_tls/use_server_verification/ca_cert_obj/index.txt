---
page_title: "webhook.http_config.use_tls.use_server_verification.ca_cert_obj"
subcategory: ""
description: "webhook.http_config.use_tls.use_server_verification.ca_cert_obj for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2336, "body_sha256": "sha256:cb44ad83faf6d4769e90abb76eacf7e389ef9d663fa1176eec5c7e5bff5d269c", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj:trusted_ca"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/index.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification", "ca_cert_obj"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.use_tls.use_server_verification.ca_cert_obj for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.use_server_verification.ca_cert_obj

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/)
- [webhook.http_config.use_tls.use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ca cert obj.

Upstream description:

Configuration for CA certificate.

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
ca_cert_obj {
  # Configure direct properties listed below.
}
```

## Direct properties

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/): complete subsection reference.

## Next pages

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/)
- [webhook.http_config.use_tls.use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
