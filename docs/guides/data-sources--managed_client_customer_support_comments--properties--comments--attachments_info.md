---
page_title: "comments.attachments_info"
subcategory: ""
description: "comments.attachments_info for xcsh_managed_client_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 1616, "body_sha256": "sha256:3cf4cdb3b1178dbfcc2ee1ebcebcc12c71d8e85f5ab977751a35d1f4aff1edcf", "canonical_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments:attachments_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments:attachments_info", "parent_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "path": "docs/guides/data-sources--managed_client_customer_support_comments--properties--comments--attachments_info.md", "provider_name": "managed_client_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["comments", "attachments_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "comments.attachments_info for xcsh_managed_client_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# comments.attachments_info

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md)
- [Property reference](data-sources--managed_client_customer_support_comments--reference.md)
- [comments](data-sources--managed_client_customer_support_comments--properties--comments.md)
- comments.attachments_info

<a id="section"></a>

Type: `"list"`. Computed.

Information about any attachments (such as screenshots, plain text files) the comment can have.

## Direct properties

<a id="schema-comments--attachments_info--attachment"></a>

### attachment property

Type: `"string"`. Computed.

Any binary attachment (such as screenshots, plain text files, PDFs) encoded as base64 if used over
HTTP.

<a id="schema-comments--attachments_info--content_type"></a>

### content_type property

Type: `"string"`. Computed.

MIME content type of the attachment. Helps the UI to properly display the data.

<a id="schema-comments--attachments_info--filename"></a>

### filename property

Type: `"string"`. Computed.

Filename of the attachment as provided by the caller.

<a id="schema-comments--attachments_info--tp_id"></a>

### tp_id property

Type: `"string"`. Computed.

Optional ID as assigned by the third-party actually storing the data.

## Next pages

- [comments](data-sources--managed_client_customer_support_comments--properties--comments.md)
- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md)
