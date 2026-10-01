---
page_title: "xcsh_managed_client_customer_support_comments reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_managed_client_customer_support_comments reference."
---

# xcsh_managed_client_customer_support_comments reference

<a id="canonical-002a4575b61eb4ad6cacc307c2f712501721ed41a3e99b73773ff4b3e8095b91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ea9407b248a540b6ea8f43c14abd2d78e7b3846acf8fc3fa0ab2802cf34ff35"></a>

## Property reference — Property reference / 6f0c0920ecbf / 2

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)
- Property reference

<a id="canonical-9c3bde16aeeff7743d52b53dc3ce52e8b944e8cf768d770cb4c75289a93d7d02"></a>

## Direct properties — Property reference / 6f0c0920ecbf / 3

<a id="canonical-fff0b809ecd0eadfa7ad73f26b8571b876fbfb2e611efeae9cd4efd1085817ba"></a>

<a id="canonical-41abae1703c1ed5f4f43daa948523e0eaee0e6f8b85f1356f76860d9bcdfc897"></a>

## all_comments_returned property — Property reference / 6f0c0920ecbf / 4

Type: `"bool"`. Computed.

Indicates if all comments for the issue have been returned.

- [comments](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-fd37fb8bc98ca6cca7212ddebd69f5ec187bc2e2ab11ab67bc71a4f9d01e3bdf): complete subsection reference.

<a id="canonical-5646a05ba81c3d4692faeeac681c95de8dfd50711793d44fc78d829052e71c70"></a>

<a id="canonical-7e61623c5afd98c86c30b86274a9fab1b399bfed2735fc8321c63e45230e6a92"></a>

## created_until_timestamp property — Property reference / 6f0c0920ecbf / 5

Type: `"string"`. Optional.

Filter to retrieve comments created up to the specified timestamp.

<a id="canonical-0ee86288c5724ba3a6181a7b907526e664c5517d294ac01d047fe68fda2e06fa"></a>

<a id="canonical-38da1585150cf7863829d3af738dda2e1cbef971638e4df858922542de298d66"></a>

## tp_id property — Property reference / 6f0c0920ecbf / 6

Type: `"string"`. Required.

Tpid ID assigned to this ticket by Third Party.

<a id="canonical-8ee2dba2db27cd378e5480352d656c0c96234ebf60f1605fbda8f1fa97f4abf8"></a>

## All schema paths — Property reference / 6f0c0920ecbf / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `all_comments_returned` | [all_comments_returned](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-fff0b809ecd0eadfa7ad73f26b8571b876fbfb2e611efeae9cd4efd1085817ba) |
| `comments` | [comments](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-8958ff75530cdc7c32479b87c88fab8662bc8a2f62bf3427b0ee7072d6d7c11b) |
| `comments.attachment_ids` | [comments.attachment_ids](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-2db7358303cfac6b7cc7e2bef14424b60b11d0485de76c0e64f11f759d0df623) |
| `comments.attachments_info` | [comments.attachments_info](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-f1a0056c44569340bd08f511ac26b3d83484c146010c24f5d7e8e75579cdd87b) |
| `comments.attachments_info.attachment` | [comments.attachments_info.attachment](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-93cb8135017d085ab58f57af1c566e1b2cfbd4883cf980a90e2fefc39e75c0da) |
| `comments.attachments_info.content_type` | [comments.attachments_info.content_type](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-58dd805566ab6ec4de3e95b9c039a72ec6fd363faa12307ae617e11275d9f890) |
| `comments.attachments_info.filename` | [comments.attachments_info.filename](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-de5a3ce8f8ae088984b6a1328483b672811e353e3e14ed0676a4835f5aa0d6ab) |
| `comments.attachments_info.tp_id` | [comments.attachments_info.tp_id](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-aa19aec03da3618e8d483fb0d85a43c59392ba32795b0f0bcae5f5875af9ba73) |
| `comments.author_email` | [comments.author_email](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-a9ce0b044b8cf1401e06d2174c5cfa33655f3890b93ebecf44811bc9c672164e) |
| `comments.author_name` | [comments.author_name](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-669268614b7f9a675eef524e6decc1c78c30b15a225618c9a86405b9f0a3189b) |
| `comments.comment_id` | [comments.comment_id](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-3c1d4de0e3239bcc905a9e668408eb23366b117b148dcf4c70d53bb81dec1b02) |
| `comments.created_at` | [comments.created_at](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-19ad5d0714784811ea3853c923e67ddad36b123580541ddc98e7c87d0152ef54) |
| `comments.html` | [comments.html](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-4e2ff23252e9638dadd70cc897a4d0b086b30a04b653ec0cc75b4a3b7a926d4a) |
| `comments.plain_text` | [comments.plain_text](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-25c46b02c75fd4edc274e9a8c546c80ae73d564bc92338d2e552e6694b745f4c) |
| `created_until_timestamp` | [created_until_timestamp](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-5646a05ba81c3d4692faeeac681c95de8dfd50711793d44fc78d829052e71c70) |
| `tp_id` | [tp_id](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-0ee86288c5724ba3a6181a7b907526e664c5517d294ac01d047fe68fda2e06fa) |

<a id="canonical-45012e0376561e28ffec01cd5a008c7df8fd6c6cde8917530773a8e29e46ab7b"></a>

## Next pages — Property reference / 6f0c0920ecbf / 8

- [comments](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-fd37fb8bc98ca6cca7212ddebd69f5ec187bc2e2ab11ab67bc71a4f9d01e3bdf)
- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)

<a id="canonical-fd37fb8bc98ca6cca7212ddebd69f5ec187bc2e2ab11ab67bc71a4f9d01e3bdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c4df66b53848d09a1fededc8cd955dc83bf60af6b638554ecde49875c2c9467"></a>

## comments — comments / 550e62abf640 / 2

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)
- [Property reference](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-002a4575b61eb4ad6cacc307c2f712501721ed41a3e99b73773ff4b3e8095b91)
- comments

<a id="canonical-8958ff75530cdc7c32479b87c88fab8662bc8a2f62bf3427b0ee7072d6d7c11b"></a>

Type: `"list"`. Computed.

List of comments on the customer support ticket.

<a id="canonical-be1e04b8cf14687dbe7da9e06cab44e5ec0ef5d81c6f9dec2a665e52278257c1"></a>

## Direct properties — comments / 550e62abf640 / 3

<a id="canonical-2db7358303cfac6b7cc7e2bef14424b60b11d0485de76c0e64f11f759d0df623"></a>

<a id="canonical-eea22d14d12c955fd7ea6f65af3ac152f5ca4b937b4a3cb6a3eb0489ddad1efd"></a>

## attachment_ids property — comments / 550e62abf640 / 4

Type: `["list", "string"]`. Computed.

Third party ID of any attachment related to this ticket comment.

- [attachments_info](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-ec60247c2dee3ffc2b1e2c29db901cafd31ad70653cea39f306d41fbde6e5828): complete subsection reference.

<a id="canonical-a9ce0b044b8cf1401e06d2174c5cfa33655f3890b93ebecf44811bc9c672164e"></a>

<a id="canonical-aa4006b8a54c01feeedc20cbc9ebfcb3b253c27bfd5be45951b759a7872d5ce8"></a>

## author_email property — comments / 550e62abf640 / 5

Type: `"string"`. Computed.

Email. Email of the author of the comment.

<a id="canonical-669268614b7f9a675eef524e6decc1c78c30b15a225618c9a86405b9f0a3189b"></a>

<a id="canonical-4ff68d0dde79a6c6d843923ad073f9eacc416bd2ce776f23ab918b429864a9d8"></a>

## author_name property — comments / 550e62abf640 / 6

Type: `"string"`. Computed.

Author. Author of the comment (as a name)

<a id="canonical-3c1d4de0e3239bcc905a9e668408eb23366b117b148dcf4c70d53bb81dec1b02"></a>

<a id="canonical-77c72a5f46fb28093b6f543693915a9bb849c13c822aa618bd5891e39adff8ab"></a>

## comment_id property — comments / 550e62abf640 / 7

Type: `"string"`. Computed.

ID assigned to this comment by support provider.

<a id="canonical-19ad5d0714784811ea3853c923e67ddad36b123580541ddc98e7c87d0152ef54"></a>

<a id="canonical-3a08579a36f5957f7dc6ac01896af724403de4516cf9f2dc7431c500c95d9e8f"></a>

## created_at property — comments / 550e62abf640 / 8

Type: `"string"`. Computed.

At. Comment creation time.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="canonical-4e2ff23252e9638dadd70cc897a4d0b086b30a04b653ec0cc75b4a3b7a926d4a"></a>

<a id="canonical-9b631fb4247dcad35f90de9806f0cde3ddebb5b21a581b47756bf5526fa23ff4"></a>

## html property — comments / 550e62abf640 / 9

Type: `"string"`. Computed.

Comment. Comment body as HTML.

<a id="canonical-25c46b02c75fd4edc274e9a8c546c80ae73d564bc92338d2e552e6694b745f4c"></a>

<a id="canonical-bd86a1275228ab0ea18d4db791fd54691b14da5942ec37d4e1e5a5d0402938cd"></a>

## plain_text property — comments / 550e62abf640 / 10

Type: `"string"`. Computed.

Comment. Comment body as plain text.

<a id="canonical-75d5a3cacaff29863fa68a8a05560fb141cd20fc355b5a1cadf9f3c8410b3ae8"></a>

## Next pages — comments / 550e62abf640 / 11

- [comments.attachments_info](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-ec60247c2dee3ffc2b1e2c29db901cafd31ad70653cea39f306d41fbde6e5828)
- [Property reference](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-002a4575b61eb4ad6cacc307c2f712501721ed41a3e99b73773ff4b3e8095b91)
- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)

<a id="canonical-ec60247c2dee3ffc2b1e2c29db901cafd31ad70653cea39f306d41fbde6e5828"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d02eac0be0be10e9d4c8470a60fd659e00505ffe71ae4802364be3d6485d916"></a>

## comments.attachments_info — comments.attachments_info / 45ffce87c736 / 2

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)
- [Property reference](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-002a4575b61eb4ad6cacc307c2f712501721ed41a3e99b73773ff4b3e8095b91)
- [comments](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-fd37fb8bc98ca6cca7212ddebd69f5ec187bc2e2ab11ab67bc71a4f9d01e3bdf)
- comments.attachments_info

<a id="canonical-f1a0056c44569340bd08f511ac26b3d83484c146010c24f5d7e8e75579cdd87b"></a>

Type: `"list"`. Computed.

Information about any attachments (such as screenshots, plain text files) the comment can have.

<a id="canonical-b8a31c206998eadaa51789ad30ba2496224404de4b891ee9ed51e32c2adfa46e"></a>

## Direct properties — comments.attachments_info / 45ffce87c736 / 3

<a id="canonical-93cb8135017d085ab58f57af1c566e1b2cfbd4883cf980a90e2fefc39e75c0da"></a>

<a id="canonical-8c5a5d8ac2b9a9679154f4c79299bbe2d0a70f611c4648476ce6ea4d6397b737"></a>

## attachment property — comments.attachments_info / 45ffce87c736 / 4

Type: `"string"`. Computed.

Any binary attachment (such as screenshots, plain text files, PDFs) encoded as base64 if used over
HTTP.

<a id="canonical-58dd805566ab6ec4de3e95b9c039a72ec6fd363faa12307ae617e11275d9f890"></a>

<a id="canonical-ad7f1fbb0058f9289af83b54001fbc3854273a4777ed19e037d5e31e936cac6e"></a>

## content_type property — comments.attachments_info / 45ffce87c736 / 5

Type: `"string"`. Computed.

MIME content type of the attachment. Helps the UI to properly display the data.

<a id="canonical-de5a3ce8f8ae088984b6a1328483b672811e353e3e14ed0676a4835f5aa0d6ab"></a>

<a id="canonical-7e4d54868e7187e4b8751f84811057951a0c29a1ffbdc4a45636ce4925a42628"></a>

## filename property — comments.attachments_info / 45ffce87c736 / 6

Type: `"string"`. Computed.

Filename of the attachment as provided by the caller.

<a id="canonical-aa19aec03da3618e8d483fb0d85a43c59392ba32795b0f0bcae5f5875af9ba73"></a>

<a id="canonical-2b19ce3bda73093930821407d161a215689283a5ad92bff8a47461830f21a587"></a>

## tp_id property — comments.attachments_info / 45ffce87c736 / 7

Type: `"string"`. Computed.

Optional ID as assigned by the third-party actually storing the data.

<a id="canonical-a8700bf4a247115f8c22f89cd8fb384fa96fd27bec18bc44357dbb112610a182"></a>

## Next pages — comments.attachments_info / 45ffce87c736 / 8

- [comments](data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-fd37fb8bc98ca6cca7212ddebd69f5ec187bc2e2ab11ab67bc71a4f9d01e3bdf)
- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)
