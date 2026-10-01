---
page_title: "do_not_advertise"
subcategory: ""
description: "do_not_advertise for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1190, "body_sha256": "sha256:2d05df615b48438575ebb3f2c9bd4ed04581aa9029a8ccde58ad626902ba50a4", "canonical_id": "xcsh-docs:resources:proxy:properties:do_not_advertise", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:do_not_advertise", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "docs/guides/resources--proxy--properties--do_not_advertise.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["do_not_advertise"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/do_not_advertise/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "do_not_advertise for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# do_not_advertise

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- do_not_advertise

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: do\_not\_advertise, site\_virtual\_sites\] Configuration parameter for do not advertise.

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

- [do_not_advertise](resources--proxy--properties--do_not_advertise.md#section)
- [site_virtual_sites](resources--proxy--properties--site_virtual_sites.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_advertise = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--proxy--reference.md)
- [xcsh_proxy](../resources/proxy.md)
