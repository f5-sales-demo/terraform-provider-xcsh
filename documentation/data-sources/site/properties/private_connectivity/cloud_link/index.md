---
page_title: "private_connectivity.cloud_link"
subcategory: "Infrastructure"
description: "private_connectivity.cloud_link for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1577, "body_sha256": "sha256:f6c138f238382aef9ac2256c2c3f85432d81a3c75339c81f9433afbb5b011785", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "parent_id": "xcsh-docs:data-sources:site:properties:private_connectivity", "path": "documentation/data-sources/site/properties/private_connectivity/cloud_link/index.md", "provider_name": "site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["private_connectivity", "cloud_link"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/private_connectivity/cloud_link/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connectivity.cloud_link for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity.cloud_link

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/private_connectivity/)
- private_connectivity.cloud_link

<a id="section"></a>

Type: `"single"`. Computed.

Information related to cloud link used by the site.

## Direct properties

<a id="schema-private_connectivity--cloud_link--name"></a>

### name property

Type: `"string"`. Computed.

Name of the the CloudLink used with this site.

<a id="schema-private_connectivity--cloud_link--state"></a>

### state property

Type: `"string"`. Computed.

\[Enum: UP|DOWN|DEGRADED|NOT\_APPLICABLE\] State of the CloudLink connections - UP: Up CloudLink and
their corresponding Direct Connect connections are up and healthy - DOWN: Down CloudLink and their
corresponding Direct Connect connections are down - DEGRADED: Degraded Some of Direct Connect
connections with the CloudLink are down .. Possible values are \`UP\`, \`DOWN\`, \`DEGRADED\`,
\`NOT\_APPLICABLE\`. Defaults to \`UP\`.

## Next pages

- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/private_connectivity/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
