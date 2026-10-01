---
page_title: "cloudflare.protected_endpoints.web_client"
subcategory: ""
description: "cloudflare.protected_endpoints.web_client for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2501, "body_sha256": "sha256:22f8010f339c4e1532250debc4742253b170844edf8840620234fe9ea6e36881", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "path": "docs/guides/resources--protected_application--properties--cloudflare--protected_endpoints--web_client.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/web_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_client for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_client

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [cloudflare.protected_endpoints](resources--protected_application--properties--cloudflare--protected_endpoints.md)
- cloudflare.protected_endpoints.web_client

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Web Client. Web client configuration OPTIONS.

Upstream description:

Web client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("continue",
    "redirect")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
web_client {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md): complete subsection reference.

- [continue](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue.md): complete subsection reference.

- [redirect](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--redirect.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_client.block](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue.md)
- [cloudflare.protected_endpoints.web_client.redirect](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--redirect.md)
- [cloudflare.protected_endpoints](resources--protected_application--properties--cloudflare--protected_endpoints.md)
- [xcsh_protected_application](../resources/protected_application.md)
