---
page_title: "comments.attachments_info"
subcategory: ""
description: "comments.attachments_info for xcsh_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 1412, "body_sha256": "sha256:11a577cb2f4b9e886592c45e27fd33e007cf699af79e1ee7fb1059db2d0c4520", "canonical_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "parent_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "path": "docs/guides/data-sources--customer_support_comments--properties--comments--attachments_info.md", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["comments", "attachments_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/properties/comments/attachments_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "comments.attachments_info for xcsh_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# comments.attachments_info

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
- [Property reference](data-sources--customer_support_comments--reference.md)
- [comments](data-sources--customer_support_comments--properties--comments.md)
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

- [comments](data-sources--customer_support_comments--properties--comments.md)
- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
