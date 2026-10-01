---
page_title: "xcsh_site_upgrade_status reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_status reference."
---

# xcsh_site_upgrade_status reference

<a id="canonical-026e1aecd2a970e096dc9476e95faa6d85dfc2ee51f7bdce20b6431348fda126"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b418747a6419b61cf00e4234e2765dac6868ab850e128aa272830db1f649baba"></a>

## Property reference — Property reference / 5849348f5c39 / 2

Breadcrumbs:

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-354cc1e18b1eae705be66b9b6e4a44fd8152a2aa13c69db8d19d96b451ff6445)
- Property reference

<a id="canonical-e4de8e0f482a5df2c9c2031dc3b827f1af5703a075c4d1764c2cdb7a2dd61807"></a>

## Direct properties — Property reference / 5849348f5c39 / 3

<a id="canonical-fb9c9a68440575af916f4020eb17961f42448186c017f2acdb8d25bff2ff196b"></a>

<a id="canonical-cd6e88c30d83293b02beab114ebb4f6e6b854051282bbed918bfdcce9ed4b32b"></a>

## eligible property — Property reference / 5849348f5c39 / 4

Type: `"bool"`. Computed.

Whether the site is ONLINE and each selected target is installed or advertised for upgrade. Software
prechecks must pass when software would change; an unchanged paired version does not block a serial
software or OS upgrade.

<a id="canonical-1919ec64635de5373555e78bb38c649a927fd7ffc0e7d46368cc13aeadb987f9"></a>

<a id="canonical-f9e90b76c99c1b9ce012392802d36a15b2643ce2c9ad942d42ad2efa3ad00ee5"></a>

## expected_os_version property — Property reference / 5849348f5c39 / 5

Type: `"string"`. Optional.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-dd2c376a9a2a70f00baa07a08768beaa2e1442a2b5cc3a102fd0cdefa745358b"></a>

<a id="canonical-7e8b7c8b4c81d89f1413c190b75255b518b6cfbda2da87bdbddffde58be3eddb"></a>

## expected_software_version property — Property reference / 5849348f5c39 / 6

Type: `"string"`. Optional.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-7b12643dc0431e45ce90263d64fcd62355b1033ce90b78f1a63145015b388d0b"></a>

<a id="canonical-95eaa8e789c706639c21778ccf908108c203ef0081643e05c23aed4d659a6a5e"></a>

## failed_precheck_names property — Property reference / 5849348f5c39 / 7

Type: `["list", "string"]`. Computed.

Failed software prechecks for a newer software target; empty when the selected software version is
already installed.

<a id="canonical-0c8781c31090f526f9e73f74a261d596da4eceeac83742b0147d21d3573a718d"></a>

<a id="canonical-6019671f54a310b6ea309f6d24d12ac90e4b026b4acda975856e6a2a0abb9736"></a>

## id property — Property reference / 5849348f5c39 / 8

Type: `"string"`. Computed.

<a id="canonical-f61ce58be63d5c08f0dd417339fa333f6222dbe7fcb984c192fdfecf87c8eb4a"></a>

<a id="canonical-4f416643635e6be6e2a67b6c215b4f7f6cde2aebc13c001f0df7a540c1bc2a91"></a>

## os_available_version property — Property reference / 5849348f5c39 / 9

Type: `"string"`. Computed.

<a id="canonical-9fd3363545b95b6b7ecd54adc5bf4e89dbc5bb28bbf9c625da0f7bdf655f2ffa"></a>

<a id="canonical-b163fc80b0e7e39a120666e0670cfd5e282621368a473c82a6c4e1661917c349"></a>

## os_deployment_phase property — Property reference / 5849348f5c39 / 10

Type: `"string"`. Computed.

<a id="canonical-0f2a8f6e04297773f9792b8083b4a43fc52828e2e0c3b104541765d28989bc2c"></a>

<a id="canonical-732d46c427e71f3ac4021a9c5dafa740247c61d6144dffc580ca2aab0af63041"></a>

## os_deployment_result property — Property reference / 5849348f5c39 / 11

Type: `"string"`. Computed.

<a id="canonical-36d9b56fbd0bfb312340d34382602a14dd0db7ffc299292b83c30c1288058e08"></a>

<a id="canonical-136fa736c7432151b61187b395e7fcd4ed8a256c0baad2c16e4da8af21246cfe"></a>

## os_installed_version property — Property reference / 5849348f5c39 / 12

Type: `"string"`. Computed.

<a id="canonical-660de612e892e2dcfe6f5f19e5e7690f1136ff11c973401cf0d988db459173f3"></a>

<a id="canonical-2cd74bdfb8da64ebd9cdab1bb0c174a35a53d475cf18a994239a1962bc35d63a"></a>

## poll_interval_seconds property — Property reference / 5849348f5c39 / 13

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 300)}
```

<a id="canonical-f08647ec409cf455908494c37008204821d1f0427d12ca0b5733d97012c4b27f"></a>

<a id="canonical-77a3397ac8072f84dbb3e09c6664786110e6252d41b7f187f517384e006d727f"></a>

## ready property — Property reference / 5849348f5c39 / 14

Type: `"bool"`. Computed.

Whether the site is operationally ready (\`ONLINE\`), independent of target eligibility.

<a id="canonical-b91bd1fb0c609b110cfe9d6fb3310911e60274961e4ab431285c96e925db2017"></a>

<a id="canonical-5232c453527ae69b3257ad91cfde62da8f5656b7586ee6a9115d3f21d4dc24f6"></a>

## site property — Property reference / 5849348f5c39 / 15

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-9841fbd176d2d3b54d390f6bb6bd1509843c0d78264e9ef5f5346815dda5fd76"></a>

<a id="canonical-0e6b6dedca794e8d2ba08536f78dd614d6fc75861678df965b2803a5b4bb06a4"></a>

## site_state property — Property reference / 5849348f5c39 / 16

Type: `"string"`. Computed.

<a id="canonical-460357a71ae26a77b78e0f872e707598915bc2ceb91bd19b42b9a895b3dfcbcd"></a>

<a id="canonical-30051ed38019ee00871cd656599d494a0088f2edecaa7456a7050eeeb4462ed8"></a>

## software_available_version property — Property reference / 5849348f5c39 / 17

Type: `"string"`. Computed.

<a id="canonical-4793e1129387f2fff42a0d25f0c34da97ae3b8e34733e34d7ee1fb0e8363ce23"></a>

<a id="canonical-69325adc5a5cbd0763a242941722b9e7fd8fb677bcc3b22630e52f6598167e25"></a>

## software_deployment_phase property — Property reference / 5849348f5c39 / 18

Type: `"string"`. Computed.

<a id="canonical-baddb73c79589ffdf31498163dfd3ab0763bcfd6f9f79ed21021a8a5d1ba94fc"></a>

<a id="canonical-c8100fee5ff27cddbb0e3a43be1ea4d51b7b350d2a0d3b72e9deda124006482e"></a>

## software_deployment_result property — Property reference / 5849348f5c39 / 19

Type: `"string"`. Computed.

<a id="canonical-c0f5a0710cba64703a3c9bcbfd89d44f79a0849dad6a4cb0aa4c356a8b32d5de"></a>

<a id="canonical-e78a0224ecfaafeef61fccdcddaedca81e4d1ab5b43b3cd82c0304fff56e018a"></a>

## software_installed_version property — Property reference / 5849348f5c39 / 20

Type: `"string"`. Computed.

<a id="canonical-40594fd2efe9813ce65ba124205e1eb2fa3b182cf0ec3597f5b375a07b28d88d"></a>

<a id="canonical-56d6f506e8378fd30db2453f04380ed8137483d76f22b1278d31733825cf6879"></a>

## target_converged property — Property reference / 5849348f5c39 / 21

Type: `"bool"`. Computed.

<a id="canonical-3c836eb8fccdced53e2a23b98b00c56217973149caf63d1829b120e48d05d84d"></a>

<a id="canonical-e01bfe10775defb6e5a48ad2895190bbd93bd15140c7fd73624bf92f98ccddd5"></a>

## timeout_seconds property — Property reference / 5849348f5c39 / 22

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 7200)}
```

<a id="canonical-0f38d9d1cb01e016f10d9ee31fdf6dbaf5d59f4ed1cf50db20b9f36a2881bf84"></a>

<a id="canonical-cfd1cb315e4d5c95832af5347011ec3cab4cf6ab5ddadba807e370c86a66cc83"></a>

## upgradable_software_versions property — Property reference / 5849348f5c39 / 23

Type: `["list", "string"]`. Computed.

<a id="canonical-0cb2676acc07d96bc51550a70cf18660a95954c0ba6c16d5210608b1d61ef486"></a>

<a id="canonical-fa43a8abe817e550dcdeba142da5ef0c8cf2bd46c77db913bbaa09c398378e9c"></a>

## wait property — Property reference / 5849348f5c39 / 24

Type: `"bool"`. Optional, Computed.

<a id="canonical-00bc5edd70b406962cfa010fd8a222c97fabd7ddee7885b29663155ac30df31a"></a>

## All schema paths — Property reference / 5849348f5c39 / 25

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `eligible` | [eligible](data-sources--site_upgrade_status--reference--group-001.md#canonical-fb9c9a68440575af916f4020eb17961f42448186c017f2acdb8d25bff2ff196b) |
| `expected_os_version` | [expected_os_version](data-sources--site_upgrade_status--reference--group-001.md#canonical-1919ec64635de5373555e78bb38c649a927fd7ffc0e7d46368cc13aeadb987f9) |
| `expected_software_version` | [expected_software_version](data-sources--site_upgrade_status--reference--group-001.md#canonical-dd2c376a9a2a70f00baa07a08768beaa2e1442a2b5cc3a102fd0cdefa745358b) |
| `failed_precheck_names` | [failed_precheck_names](data-sources--site_upgrade_status--reference--group-001.md#canonical-7b12643dc0431e45ce90263d64fcd62355b1033ce90b78f1a63145015b388d0b) |
| `id` | [id](data-sources--site_upgrade_status--reference--group-001.md#canonical-0c8781c31090f526f9e73f74a261d596da4eceeac83742b0147d21d3573a718d) |
| `os_available_version` | [os_available_version](data-sources--site_upgrade_status--reference--group-001.md#canonical-f61ce58be63d5c08f0dd417339fa333f6222dbe7fcb984c192fdfecf87c8eb4a) |
| `os_deployment_phase` | [os_deployment_phase](data-sources--site_upgrade_status--reference--group-001.md#canonical-9fd3363545b95b6b7ecd54adc5bf4e89dbc5bb28bbf9c625da0f7bdf655f2ffa) |
| `os_deployment_result` | [os_deployment_result](data-sources--site_upgrade_status--reference--group-001.md#canonical-0f2a8f6e04297773f9792b8083b4a43fc52828e2e0c3b104541765d28989bc2c) |
| `os_installed_version` | [os_installed_version](data-sources--site_upgrade_status--reference--group-001.md#canonical-36d9b56fbd0bfb312340d34382602a14dd0db7ffc299292b83c30c1288058e08) |
| `poll_interval_seconds` | [poll_interval_seconds](data-sources--site_upgrade_status--reference--group-001.md#canonical-660de612e892e2dcfe6f5f19e5e7690f1136ff11c973401cf0d988db459173f3) |
| `ready` | [ready](data-sources--site_upgrade_status--reference--group-001.md#canonical-f08647ec409cf455908494c37008204821d1f0427d12ca0b5733d97012c4b27f) |
| `site` | [site](data-sources--site_upgrade_status--reference--group-001.md#canonical-b91bd1fb0c609b110cfe9d6fb3310911e60274961e4ab431285c96e925db2017) |
| `site_state` | [site_state](data-sources--site_upgrade_status--reference--group-001.md#canonical-9841fbd176d2d3b54d390f6bb6bd1509843c0d78264e9ef5f5346815dda5fd76) |
| `software_available_version` | [software_available_version](data-sources--site_upgrade_status--reference--group-001.md#canonical-460357a71ae26a77b78e0f872e707598915bc2ceb91bd19b42b9a895b3dfcbcd) |
| `software_deployment_phase` | [software_deployment_phase](data-sources--site_upgrade_status--reference--group-001.md#canonical-4793e1129387f2fff42a0d25f0c34da97ae3b8e34733e34d7ee1fb0e8363ce23) |
| `software_deployment_result` | [software_deployment_result](data-sources--site_upgrade_status--reference--group-001.md#canonical-baddb73c79589ffdf31498163dfd3ab0763bcfd6f9f79ed21021a8a5d1ba94fc) |
| `software_installed_version` | [software_installed_version](data-sources--site_upgrade_status--reference--group-001.md#canonical-c0f5a0710cba64703a3c9bcbfd89d44f79a0849dad6a4cb0aa4c356a8b32d5de) |
| `target_converged` | [target_converged](data-sources--site_upgrade_status--reference--group-001.md#canonical-40594fd2efe9813ce65ba124205e1eb2fa3b182cf0ec3597f5b375a07b28d88d) |
| `timeout_seconds` | [timeout_seconds](data-sources--site_upgrade_status--reference--group-001.md#canonical-3c836eb8fccdced53e2a23b98b00c56217973149caf63d1829b120e48d05d84d) |
| `upgradable_software_versions` | [upgradable_software_versions](data-sources--site_upgrade_status--reference--group-001.md#canonical-0f38d9d1cb01e016f10d9ee31fdf6dbaf5d59f4ed1cf50db20b9f36a2881bf84) |
| `wait` | [wait](data-sources--site_upgrade_status--reference--group-001.md#canonical-0cb2676acc07d96bc51550a70cf18660a95954c0ba6c16d5210608b1d61ef486) |

<a id="canonical-12996f4f92152dbd922d978ba8d57a1d30b7af4e13163a78a18841194eba773c"></a>

## Next pages — Property reference / 5849348f5c39 / 26

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-354cc1e18b1eae705be66b9b6e4a44fd8152a2aa13c69db8d19d96b451ff6445)
