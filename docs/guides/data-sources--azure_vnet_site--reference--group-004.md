---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-1102210000220203-1230010112220122-2102131103313233-2003033200131030-2000002203020221-2331301320111031-3323100331322331-2302302123101301"></a>

## Next pages — connections / 110223210132 / 6

- [ingress_egress_gw.hub.express_route_enabled.connections.metadata](data-sources--azure_vnet_site--reference--group-004.md#canonical-2000031132022130-0232312223300213-3300021103031310-2101231103013000-1310031212320210-2330230013023132-3312003203201102-3223222203122232)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-004.md#canonical-3332012222230133-3130120110213202-1313310311231113-1110220312302133-1003213110202311-2330022202320231-2130210103203111-1333132030222131)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2000031132022130-0232312223300213-3300021103031310-2101231103013000-1310031212320210-2330230013023132-3312003203201102-3223222203122232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111202011110033-0230012012002111-1110031322232003-3001321111213012-0122231033220101-0330230001202123-2130200112011220-1221331101322200"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.metadata — metadata / 122003302030 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-3210002321332310-0011030010113313-3332113011202123-0303211101102010-0312031312031003-2012001301300201-0020100111200132-3120322210101200)
- ingress_egress_gw.hub.express_route_enabled.connections.metadata

<a id="canonical-3333312300300233-3023010022000030-2031010311113021-1333222112323323-1331230122103111-1310220303031010-3132333232310211-3003003003100012"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2022313213021332-3332130012220011-2111012211201120-1122013212230233-0002022322011230-2020003131201330-3300312102211233-3032212023102220"></a>

## Direct properties — metadata / 122003302030 / 3

<a id="canonical-3102022231302331-2113202302013030-3001112320331110-0001111022033200-1113112013303001-0020031000222203-1000012112012202-2103331122002301"></a>

<a id="canonical-2313112110001020-2031332133323233-1130112221113203-3010120332323212-0232022201111213-2210312130321033-1231321100021223-1232001031301121"></a>

## description_spec property — metadata / 122003302030 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1201320113102202-0232121323020232-1303322011321211-1321010020133313-1100023210223312-1000133003312123-3222320010211312-0230303332301320"></a>

<a id="canonical-2033032332200220-3221012101103233-1223012301111320-3330000330322210-1332022322313200-1330230320130130-3101333202010320-0031202332123322"></a>

## name property — metadata / 122003302030 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1312031011221221-3210211010232112-0133001011110222-2010102011111000-0132220231302133-2031023231131303-1312211020122002-0201120233220321"></a>

## Next pages — metadata / 122003302030 / 6

- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-3210002321332310-0011030010113313-3332113011202123-0303211101102010-0312031312031003-2012001301300201-0020100111200132-3120322210101200)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3332012222230133-3130120110213202-1313310311231113-1110220312302133-1003213110202311-2330022202320231-2130210103203111-1333132030222131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313330030300203-0113011122312103-1200033031212222-3333000312023330-2200010123221332-3002303132331003-1021003121112222-3012332032113310"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription — other_subscription / 320312110333 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-3210002321332310-0011030010113313-3332113011202123-0303211101102010-0312031312031003-2012001301300201-0020100111200132-3120322210101200)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription

<a id="canonical-0111220310102010-3233110330002101-0333011002212021-3302120122320323-0301231221122221-2212012002211022-3120132030031023-3213210000013121"></a>

Type: `"single"`. Computed.

Express Route Circuit Config From Other Subscription.

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

<a id="canonical-1332132122322101-2131011311120132-0212323021201232-3012303203112200-0320003313331311-1223003101112110-1220033103323033-1133033330331332"></a>

## Direct properties — other_subscription / 320312110333 / 3

- [authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-2110203131321023-2110030201312032-3201221220322131-0110111331121001-3332213002000202-1020210000210222-3322201003131013-0020213232011022): complete subsection reference.

<a id="canonical-2123001033211020-3331203112121223-1222202332322010-1103231100031133-1003211331332311-1110300101100000-0313010213021012-1121320311023021"></a>

<a id="canonical-3100013212121203-3231323302230300-3131112023122320-3101331232312231-0331310311302112-0310302301212023-2301010231101030-0003010210300011"></a>

## circuit_id property — other_subscription / 320312110333 / 4

Type: `"string"`. Computed.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-0100302203300323-2003021200111010-3121131201212003-2303030233130033-0332223002223100-1111332121223211-3202033110110122-3031201120311100"></a>

## Next pages — other_subscription / 320312110333 / 5

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-2110203131321023-2110030201312032-3201221220322131-0110111331121001-3332213002000202-1020210000210222-3322201003131013-0020213232011022)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-3210002321332310-0011030010113313-3332113011202123-0303211101102010-0312031312031003-2012001301300201-0020100111200132-3120322210101200)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2110203131321023-2110030201312032-3201221220322131-0110111331121001-3332213002000202-1020210000210222-3322201003131013-0020213232011022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211232002112200-3133203012232331-2310323131011002-1031202113332322-1131202220322103-1200132320010022-2000000031010222-3131031002122101"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key — authorized_key / 203223101222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-3210002321332310-0011030010113313-3332113011202123-0303211101102010-0312031312031003-2012001301300201-0020100111200132-3120322210101200)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-004.md#canonical-3332012222230133-3130120110213202-1313310311231113-1110220312302133-1003213110202311-2330022202320231-2130210103203111-1333132030222131)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="canonical-0122333231333003-1123300022102331-2012023100102012-0010320310212333-0131111331022033-3023101213010022-1323000220332301-1122323123121322"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-2101011212333012-2111231012211320-1303101023101223-3300111021111330-1130211011023203-3231033011021223-1310211001101222-0123332112131332"></a>

## Direct properties — authorized_key / 203223101222 / 3

- [blindfold_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-0202331003323120-2100000210332312-1233133233301202-2320022023022230-1231122130012032-1020202132112311-3132323120223100-2103011222130313): complete subsection reference.

- [clear_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-3103310200323121-3232302200213103-3300312231221122-1033222010131122-3022033001211211-3220013022232230-3231110131213301-1030033031313102): complete subsection reference.

<a id="canonical-3322122323003133-2201210000021032-1211022132022132-2331110023003022-3003123222001201-3113023000011200-1302010021320002-1223131023011123"></a>

## Next pages — authorized_key / 203223101222 / 4

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-0202331003323120-2100000210332312-1233133233301202-2320022023022230-1231122130012032-1020202132112311-3132323120223100-2103011222130313)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-3103310200323121-3232302200213103-3300312231221122-1033222010131122-3022033001211211-3220013022232230-3231110131213301-1030033031313102)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-004.md#canonical-3332012222230133-3130120110213202-1313310311231113-1110220312302133-1003213110202311-2330022202320231-2130210103203111-1333132030222131)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0202331003323120-2100000210332312-1233133233301202-2320022023022230-1231122130012032-1020202132112311-3132323120223100-2103011222130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210331322122232-1020312320101232-1032012013121332-2220222002333113-2032301031221033-3100021111221232-1302331112111000-0031110221212330"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info — blindfold_secret_info / 120111102131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-3210002321332310-0011030010113313-3332113011202123-0303211101102010-0312031312031003-2012001301300201-0020100111200132-3120322210101200)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-004.md#canonical-3332012222230133-3130120110213202-1313310311231113-1110220312302133-1003213110202311-2330022202320231-2130210103203111-1333132030222131)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-2110203131321023-2110030201312032-3201221220322131-0110111331121001-3332213002000202-1020210000210222-3322201003131013-0020213232011022)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="canonical-0232211101233101-2233112211331330-3231011213310213-0032010211310210-2223320020321131-0310233323113121-1031112123311330-3032301030030103"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-1331330011320321-1102110303221201-0302330211032033-3321210002333200-1011232333210022-0032320313031233-0023303311302300-1301002303230331"></a>

## Direct properties — blindfold_secret_info / 120111102131 / 3

<a id="canonical-2000222132100230-1122010312312013-2232321200311030-1130102120320022-3220223112020311-0031213122310032-0212113000212201-3233132031302103"></a>

<a id="canonical-0001312012300303-0221013300303110-0030111303201220-3320322112100112-0303203233023102-2310201213131311-2111230102000231-2220101001101322"></a>

## decryption_provider property — blindfold_secret_info / 120111102131 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0112333120233012-1101333102312313-2112301220123021-3012023212223202-0003003123113021-2202112201202333-2333302102000032-0322002013013331"></a>

<a id="canonical-0331223232330113-0322012100113232-0312131003001232-3032210021011122-2333033120230130-1232323311203012-1123212001211133-2200121320331223"></a>

## location property — blindfold_secret_info / 120111102131 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3301012232030201-0330021130111230-1202233021310302-1000233332300031-2112332203030111-3312313130002110-1103200003133320-1300301210112033"></a>

<a id="canonical-2132122311013130-0201213133302323-0312223300000131-3110002133213002-1210221003333022-0311202131131020-2130001313331323-0023000203022001"></a>

## store_provider property — blindfold_secret_info / 120111102131 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1113002102201031-2301320010100332-2311323223320201-3303120030112332-1031313123330101-1232023023130312-3132233311020312-2021313023110003"></a>

## Next pages — blindfold_secret_info / 120111102131 / 7

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-2110203131321023-2110030201312032-3201221220322131-0110111331121001-3332213002000202-1020210000210222-3322201003131013-0020213232011022)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3103310200323121-3232302200213103-3300312231221122-1033222010131122-3022033001211211-3220013022232230-3231110131213301-1030033031313102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030231000320333-2322333110211320-3203113213300200-3130100130123333-1031020121332022-0021333333011020-2320103332202211-3320012231330102"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info — clear_secret_info / 313023222300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-3210002321332310-0011030010113313-3332113011202123-0303211101102010-0312031312031003-2012001301300201-0020100111200132-3120322210101200)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-004.md#canonical-3332012222230133-3130120110213202-1313310311231113-1110220312302133-1003213110202311-2330022202320231-2130210103203111-1333132030222131)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-2110203131321023-2110030201312032-3201221220322131-0110111331121001-3332213002000202-1020210000210222-3322201003131013-0020213232011022)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info

<a id="canonical-1212012123103110-3103131302123130-2010313230320010-1211321332220031-0102012022133023-1111303313122300-0030212200111320-2110321321031001"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-1003113331323103-1202302203020102-3221201333333133-1133033132032203-2102101033211131-3003331102111202-3211022233100033-1222131212121123"></a>

## Direct properties — clear_secret_info / 313023222300 / 3

<a id="canonical-3023300311313012-0323211120130223-3132033233212310-3023212321212321-1010222132101020-1032233300302312-0312301100200111-3223000333303020"></a>

<a id="canonical-0002031112131330-1331302300203211-2201213021211302-3020232303232003-0302303110113003-1130330013211133-0321202310310013-2311121101111232"></a>

## provider_ref property — clear_secret_info / 313023222300 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0132020232220002-0002212031233132-3303132031132312-0022231122221230-3130332101013210-3120122213002133-3000101200013220-3321122000100111"></a>

<a id="canonical-0110331203302330-1101230233222131-2123220330233030-3230201023303120-1331210200103211-2110000332220132-0103010310013202-0303101123110332"></a>

## URL property — clear_secret_info / 313023222300 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2322113221330200-0110323121032031-0231233000100001-0323201031322110-0311020103211001-2110200122020113-3313030303023000-2112233032123123"></a>

## Next pages — clear_secret_info / 313023222300 / 6

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-2110203131321023-2110030201312032-3201221220322131-0110111331121001-3332213002000202-1020210000210222-3322201003131013-0020213232011022)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3333301310221211-1323103201222330-3133010123030003-0111022211002301-0111201310233021-2120203131120101-3313230331300200-1300222021222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120223202210001-1022223300022113-2011001021113333-2102000333112200-3100100132031213-2311301001200323-3023222222030130-0330231202313101"></a>

## ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server — do_not_advertise_to_route_server / 211112221103 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server

<a id="canonical-3300000101110033-0210233201011110-3233003311333212-0033123110331301-2030203031020033-3312231013300322-1121320103021231-1101021320311230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise to route server.

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

<a id="canonical-1300021033120333-1023211012001010-0120201332301212-1212030202213202-0331023223101323-2201213322110102-0310322333131211-0020320233231302"></a>

## Direct properties — do_not_advertise_to_route_server / 211112221103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310201201200110-3233232210302012-3130312032210230-0302103110211221-1223033300313130-1122233120332021-3012012333133230-0010010310102122"></a>

## Next pages — do_not_advertise_to_route_server / 211112221103 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113302200013012-3203200301210213-2031130130023212-1120202323131202-0223101320333221-0320103333011203-0332321030131110-1202230322122023"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet — gateway_subnet / 231322013021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet

<a id="canonical-2122322331112300-1101311111233330-1031100032110021-2222302232302133-1013101012101201-1300001333000202-3022013013200111-1002223120231320"></a>

Type: `"single"`. Computed.

Configuration parameter for gateway subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-0003033122301020-2131311010113221-2200223003233200-2333332132122030-1320002021221010-0131122120100001-2333113122212013-0213302211203010"></a>

## Direct properties — gateway_subnet / 231322013021 / 3

- [auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111100113232331-1321312121110010-1332221201133120-1123102112320112-1330110323202310-1122333001110011-3311232012112032-1003130123301321): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1321032011330011-1212330013230330-1133010332030202-0113211103200211-1132101022032301-0301231312132310-1002103231010202-0031201021333102): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-0331230322120130-1033202220223002-1333221110201102-2023202222112131-2112311220121013-3222102202330302-0000122301003023-1200230122330012): complete subsection reference.

<a id="canonical-2123302230312231-3333203313312213-2233123120223031-1112002022330322-3303000323211231-2212310200231123-2032310131013111-2033202121033003"></a>

## Next pages — gateway_subnet / 231322013021 / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111100113232331-1321312121110010-1332221201133120-1123102112320112-1330110323202310-1122333001110011-3311232012112032-1003130123301321)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1321032011330011-1212330013230330-1133010332030202-0113211103200211-1132101022032301-0301231312132310-1002103231010202-0031201021333102)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-0331230322120130-1033202220223002-1333221110201102-2023202222112131-2112311220121013-3222102202330302-0000122301003023-1200230122330012)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1111100113232331-1321312121110010-1332221201133120-1123102112320112-1330110323202310-1122333001110011-3311232012112032-1003130123301321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033002033203110-3323231130301003-1300211221101331-1323231100220133-1101212010312111-0111122321200200-2012100023101212-1203001110203213"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto — auto / 320101010222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto

<a id="canonical-3331130322300311-1001310333022201-1001000333212332-1312303300023222-0001232021201112-3230100223322001-1300310000210220-3123003111112131"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-2012232231310032-1001032100120101-2230301213322033-2003213112332232-3130323210022302-3201031332331330-3230000233022131-0310121301113332"></a>

## Direct properties — auto / 320101010222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113133202210123-0020232010000131-1223123231212120-1001330223020330-0302103201010111-3130222123010331-0032022203322033-1230102310003311"></a>

## Next pages — auto / 320101010222 / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1321032011330011-1212330013230330-1133010332030202-0113211103200211-1132101022032301-0301231312132310-1002103231010202-0031201021333102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312203133313332-2200230023313233-0302333123011201-1303132333133102-3223032302101323-0201212313321323-1303122131121223-1111210123022010"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet — subnet / 122202323030 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet

<a id="canonical-2111102210101121-2201220133331333-1233221123101011-2021112101010221-1110132001122001-3210113300323313-1110310113133332-0311002322113301"></a>

Type: `"single"`. Computed.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-0132121123011320-2321201112032022-0022022112200022-0013022323132320-0221201113003121-2133003121230000-1312023312133213-1122013212103221"></a>

## Direct properties — subnet / 122202323030 / 3

<a id="canonical-0302011011001323-2012031212233322-2301203113223222-1200111200323132-2101302220111222-3313300320032010-1103233133032203-0330011133102202"></a>

<a id="canonical-3203223213003101-1233122121321132-3123100011320011-0212030132311123-0012312021013200-0113001332201323-0323203332201320-3020112310332231"></a>

## subnet_resource_grp property — subnet / 122202323030 / 4

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-3030122220132020-1030133102013212-0201210030231311-1231312232121133-3301032330230323-0011303213311201-0010123000321203-0131332312020111): complete subsection reference.

<a id="canonical-0300103022012232-1230023322110030-0212230110212212-3321123021200203-0332110101310020-0212122331230131-0030303311201003-3101002312003311"></a>

## Next pages — subnet / 122202323030 / 5

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-3030122220132020-1030133102013212-0201210030231311-1231312232121133-3301032330230323-0011303213311201-0010123000321203-0131332312020111)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3030122220132020-1030133102013212-0201210030231311-1231312232121133-3301032330230323-0011303213311201-0010123000321203-0131332312020111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123031201301332-0112311311212002-0310311212231212-3303201022120230-2303000320303112-1203200230222113-0212013210012101-1202111130300302"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group — vnet_resource_group / 222210231230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1321032011330011-1212330013230330-1133010332030202-0113211103200211-1132101022032301-0301231312132310-1002103231010202-0031201021333102)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

<a id="canonical-3232110032323111-1101331313210300-2100021313212330-2202202220320030-2330032101102021-2002333021000032-3000033323123223-2130320123322313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-3222012310000203-2103332333210223-1121302321133222-1121013200311311-3121123012301302-2131033320232312-0312102032122131-3332111320013200"></a>

## Direct properties — vnet_resource_group / 222210231230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200203110221013-3002221303300312-1022332021031300-3012221101210030-1310120032111333-3321122020111230-0021331103222011-2212122002220202"></a>

## Next pages — vnet_resource_group / 222210231230 / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1321032011330011-1212330013230330-1133010332030202-0113211103200211-1132101022032301-0301231312132310-1002103231010202-0031201021333102)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0331230322120130-1033202220223002-1333221110201102-2023202222112131-2112311220121013-3222102202330302-0000122301003023-1200230122330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133110233032202-1301021221301020-1002012131010313-3203202033201310-3032303201111202-0010132000003011-0310020001302203-3013303133130032"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param — subnet_param / 123330331302 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param

<a id="canonical-3130022230201001-2320112121002220-0231201302322012-2231302221212010-1211330102012131-2232101211213120-0132103031223021-1132302003121222"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-1003312223000201-0013010302333213-0232031102031100-0212301213212323-1022131330311122-1112023223203011-0110301331031222-3200020310231020"></a>

## Direct properties — subnet_param / 123330331302 / 3

<a id="canonical-0003112132122211-3120102330233201-1212211011210231-0332100221201322-0331123103113011-1020102233122102-3330122010313221-1220001300000313"></a>

<a id="canonical-1120001322322230-2030231332200012-1332200200033003-1303333000302330-2133002101232011-0011121213130212-2101302032120303-3311312313113022"></a>

## IPv4 property — subnet_param / 123330331302 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-2122103102302231-0230223033003331-1000310210013101-1302221311333110-3020301123012020-2113121132302212-0223013332203103-0232000300132021"></a>

## Next pages — subnet_param / 123330331302 / 5

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0030123232223112-2213232033020330-1300212320300122-1233011212302223-1000120032112022-0132321201200202-3122202222023201-2113030321332133)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203122132030211-2333100222020333-1101000210312320-3211011313322102-0122002121133121-2110310022003103-0323002322110323-1213110003220030"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet — route_server_subnet / 320200120031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet

<a id="canonical-0112130011301021-0112122220302300-3100330301330232-0312112202102233-2112220013331312-3331111110210030-3211320002220310-3121202023333232"></a>

Type: `"single"`. Computed.

Configuration parameter for route server subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-3121032102023033-0100300301301302-3300023231220011-1202321021112100-2101321031023011-2301303313113013-3001133100320030-2221220301203203"></a>

## Direct properties — route_server_subnet / 320200120031 / 3

- [auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-1021300000133301-2012012123231200-1202201030102223-0011320302123012-2233320212233023-0013002022032313-3323001301232110-0123011100012101): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0122010121030301-3203332001021302-2312000123021112-3321010111110330-0112111123233333-2000303203121232-0220223033100221-0133030232032023): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-1220021031021200-1002200201100200-0010120132303113-0233333033231131-1020221302220021-3103230112012131-2323123013231131-0303310111233220): complete subsection reference.

<a id="canonical-1300233012110131-1110112101121013-3332133001032033-3021303013123323-2123210011033210-2213302120212131-2231332130200223-3202022333221203"></a>

## Next pages — route_server_subnet / 320200120031 / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-1021300000133301-2012012123231200-1202201030102223-0011320302123012-2233320212233023-0013002022032313-3323001301232110-0123011100012101)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0122010121030301-3203332001021302-2312000123021112-3321010111110330-0112111123233333-2000303203121232-0220223033100221-0133030232032023)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-1220021031021200-1002200201100200-0010120132303113-0233333033231131-1020221302220021-3103230112012131-2323123013231131-0303310111233220)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1021300000133301-2012012123231200-1202201030102223-0011320302123012-2233320212233023-0013002022032313-3323001301232110-0123011100012101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232222111002312-3122032201030211-0230011311002002-3303010010023230-2101102012013122-3223330221132303-3313021332220121-1011112011120013"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto — auto / 300322032123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto

<a id="canonical-0102222203133303-0333000033111320-0013202202230231-3210002021012313-1132233111003012-2210123033313011-2000203102313332-1331202022021130"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-3322132130302020-1333211320201213-3313112021123220-0023101022203131-3333020221200031-2121321221030103-0220311011130311-0020200111332110"></a>

## Direct properties — auto / 300322032123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011121231322300-2231103001113003-2333021131023133-2302121011323000-0200023123203332-0213013130013300-1211020323230032-3132210110222020"></a>

## Next pages — auto / 300322032123 / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0122010121030301-3203332001021302-2312000123021112-3321010111110330-0112111123233333-2000303203121232-0220223033100221-0133030232032023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330102220330321-0012031002313002-1303311332332331-2221233033103101-2323330111300311-3332001322012300-3200031200100311-0022312233022103"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet — subnet / 311203012100 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet

<a id="canonical-0212220203331303-3010002100233122-0323010013103232-1302113310103122-2112031333130233-2213221200103022-2311113322013011-2222000320021021"></a>

Type: `"single"`. Computed.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-0302121031210000-1010121002012311-2131211323310202-2301322310030100-2113003033032223-3030300103013303-1023100320210202-0101121210000210"></a>

## Direct properties — subnet / 311203012100 / 3

<a id="canonical-2001321000021310-2110020032023023-0302130302200311-1333233030110110-1131013321313321-2221200220333302-3113202200011001-1211010113322301"></a>

<a id="canonical-3012333022221233-0003221031300213-3120002220232132-2230331302313132-1013231233320333-0101112302113033-2322032011123322-3103312220103110"></a>

## subnet_resource_grp property — subnet / 311203012100 / 4

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-1323133100000033-0213011022323003-2101320302131021-1033313011012201-2010313032131131-0333202211130330-1213221213011102-0301313211110020): complete subsection reference.

<a id="canonical-2032110100023002-1213220033011231-1122203112101002-1130001013300013-2230331311113210-0113013223231100-1131301330131313-3233313333120323"></a>

## Next pages — subnet / 311203012100 / 5

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-1323133100000033-0213011022323003-2101320302131021-1033313011012201-2010313032131131-0333202211130330-1213221213011102-0301313211110020)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1323133100000033-0213011022323003-2101320302131021-1033313011012201-2010313032131131-0333202211130330-1213221213011102-0301313211110020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202220312232332-2033213012221313-3311130301302202-3020102131310232-0011313113020131-0033102101013322-1331302131112123-0103232212123323"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group — vnet_resource_group / 313111131113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0122010121030301-3203332001021302-2312000123021112-3321010111110330-0112111123233333-2000303203121232-0220223033100221-0133030232032023)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group

<a id="canonical-2111021331213201-3321303002131202-1012223310110311-2131131323111201-1112202020222220-3211000023302100-2031123332000201-2030033223032232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-1232001113102203-0033311001011130-2320033133223120-0022312211110311-2322311233231120-2021332231001311-3023011223221231-3311133330303132"></a>

## Direct properties — vnet_resource_group / 313111131113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231103000302233-1000110031013020-3330330232112021-3312313332131113-3321221131220100-2230123132310201-1300213231300232-0022233331001332"></a>

## Next pages — vnet_resource_group / 313111131113 / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0122010121030301-3203332001021302-2312000123021112-3321010111110330-0112111123233333-2000303203121232-0220223033100221-0133030232032023)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1220021031021200-1002200201100200-0010120132303113-0233333033231131-1020221302220021-3103230112012131-2323123013231131-0303310111233220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302201322112303-3210103301012122-0000013321332023-1322230100303002-2021301103123213-1033113212233111-0032123311013203-1100112223322110"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param — subnet_param / 333203120332 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param

<a id="canonical-0110331300031103-0120221230321302-3103003220202111-2121111011321022-2102111331211331-0313101213222031-0011313300020132-1022202313000233"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-1213033210010023-1130030102022202-0031301301003031-3333132303120102-0200123220310122-3121212221110303-1221101321130023-0030121111120233"></a>

## Direct properties — subnet_param / 333203120332 / 3

<a id="canonical-3033233113033333-1020211022223312-2101113303301220-0301320322321320-1313021331122103-2030022001221113-3312020001122101-0132032123122021"></a>

<a id="canonical-2201232010010013-2321230330303303-2331203130231223-1032133321221210-1222331022322233-0022103301100313-1110211022011300-1232030212031033"></a>

## IPv4 property — subnet_param / 333203120332 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-0023221230130112-2123031112330302-3122331220323323-3323310202313210-2311111132032310-0233102333122132-3200023101121112-2232132201330032"></a>

## Next pages — subnet_param / 333203120332 / 5

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1320201111103110-0332212031322320-2303203033201320-1302332332011000-1130010113010333-3101031230131333-1013110201020123-3022122213100110)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2103131232201130-0021123131101203-0103131010110002-1322311222032012-2212220331213021-1000333211331321-0303033012332332-1110230111112113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332123113230310-2333323122300321-3113133233111121-2223133332013003-3131203003031012-0130012312123313-2022011121323230-1311003102211100"></a>

## ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route — site_registration_over_express_route / 333223312221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route

<a id="canonical-2202032211210222-3100233100021033-2032002222323020-0111221323311313-1322020013002011-0311110103213211-1123020300232313-2012010102100102"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

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

<a id="canonical-1233233223203230-3302202223001220-3001120221102122-1202232132320300-2211332321000123-1310011123212110-1032032202221000-1200103000033233"></a>

## Direct properties — site_registration_over_express_route / 333223312221 / 3

<a id="canonical-1110023113022322-2123011121232111-2110223302330011-1111131200111003-0221200122033002-0223102200320332-3031320021332132-3210202302123032"></a>

<a id="canonical-2112221013131001-0213031311011313-2101121101323022-1302201112133230-0030030003232033-2323100201021321-1300010331200033-0110313301003031"></a>

## cloudlink_network_name property — site_registration_over_express_route / 333223312221 / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2133312210333223-2003221301322021-3222301003102210-0102133231032202-0301321301121230-0302310300121100-1120102020121233-0013221201012210"></a>

## Next pages — site_registration_over_express_route / 333223312221 / 5

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2132102311013213-0030001213023101-1302313111212232-0133033023301212-2033122233200303-2100110333321133-0120023010000212-0232300123320133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011023123130233-0332002023232113-3003011013120303-3031312111011012-2222320122230220-1332001113032002-1101031302012012-3013302121133022"></a>

## ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet — site_registration_over_internet / 001001133022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet

<a id="canonical-2220202301311022-1203120121131103-3300310212301320-2222100323013232-0020320210033023-0213012102210333-3110211001301001-1122002001313201"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-2130221331103230-1312011203301012-2313132210002310-1200022211110000-1320223211033001-3330322211032322-0301231330220121-1111321112102100"></a>

## Direct properties — site_registration_over_internet / 001001133022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130330220331110-1302132123011111-2320313111220300-1031001222033231-1321232033233032-3003001203222002-0031322123322220-1101223123321223"></a>

## Next pages — site_registration_over_internet / 001001133022 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3130331000310233-0200311321133002-2003003311202200-0030332312023113-1313021223202021-0332012210002112-3033011322332130-1231110112310003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232103011233203-2320033321301023-3220210313132113-1000212201313210-3133201313011002-3203001121132032-0130322130320200-0113222101030221"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_ergw1az — sku_ergw1az / 013330311012 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.sku_ergw1az

<a id="canonical-3311300221122312-2110233110022213-3231023201233310-2331031230211321-0132101301200312-0101210000030202-3231322201022010-2123210313103110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku ergw1az.

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

<a id="canonical-1230320110123011-1222313310003032-3132203230222221-2111121231212330-0222221312333202-3023011031102333-0131030112321021-0333110331232212"></a>

## Direct properties — sku_ergw1az / 013330311012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101033202200033-1211223000202133-3130121222010213-2111223323111230-1211321120201232-2121333133320323-2020212233310012-3212333110320123"></a>

## Next pages — sku_ergw1az / 013330311012 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0003311130322330-3021111102121311-2221030202012233-0230012332331221-3211000121121312-0302030332020231-3131122030020002-3123022132223002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001233113001302-0010312203321112-1301011332231130-0202133030111322-3133212121330121-1312233003122110-2313301021021021-1031223132012030"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_ergw2az — sku_ergw2az / 001100212023 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.sku_ergw2az

<a id="canonical-0010301132303122-1302110013032222-0313211111303023-2303131300320021-3012303230311001-3201312332122201-0303302232231320-1213100113332300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku ergw2az.

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

<a id="canonical-2200102332233330-0030211132312222-2103323322102303-0133231210301203-0322301033223203-0032220133100123-0020332011223311-0001022110012110"></a>

## Direct properties — sku_ergw2az / 001100212023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333231323121111-0020103131123033-1032223222212130-0331233220230223-1020211132320130-2231011203232023-3330011231121202-0100333020312311"></a>

## Next pages — sku_ergw2az / 001100212023 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0012211213330213-0132202123313133-1110310300330221-1232003030232330-3123132302120033-3103132323301233-3003331210123321-1022323303000111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132332321300132-0313102100131130-2330313131102300-1311001330021323-1230022012030333-0023221130032110-2302122012200130-1133331023000200"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_high_perf — sku_high_perf / 113031102220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.sku_high_perf

<a id="canonical-0122233302313301-2201111202000123-3232111031202231-2321202032003230-3233320101033321-2003030012221013-1021302313012102-3100233221203112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku high perf.

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

<a id="canonical-2222302300013222-3000330032223110-1031020203223122-0010230212010002-0003330233330200-0310002203132002-2033233032022302-0300132320011211"></a>

## Direct properties — sku_high_perf / 113031102220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011320111111101-3013013003010022-1222210023212202-3011022012101321-3333100101001331-0032210213213232-1333121023203200-2310200013221222"></a>

## Next pages — sku_high_perf / 113031102220 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1002332113300123-0203122130032321-0222021103030011-0312031002320021-0322101301210002-2033311030130300-1230100020212301-0120111220031020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003031032002303-3133301032320111-3330310113110100-3112010033331130-2130201011023201-2112130022012132-0222031102003010-3213300113301000"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_standard — sku_standard / 322312001213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- ingress_egress_gw.hub.express_route_enabled.sku_standard

<a id="canonical-2011110011312001-0112113102323123-1223322330120211-2132031312202130-2301310203312202-1003210323012133-1330000302231122-2011332013200112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku standard.

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

<a id="canonical-1311020002122000-1121203212232122-1102210220013200-3203232022022310-3211322231013223-2302021000310102-2103232102313320-1211123201313111"></a>

## Direct properties — sku_standard / 322312001213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202000220031110-0032210332231012-2131331321203202-3111130313111101-2123000121311303-3300202012021231-3032031120012131-2121232222132030"></a>

## Next pages — sku_standard / 322312001213 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-3113331323320313-0330233323023320-0111321203311003-1200220101311002-1002311332322001-0322123301322022-0200021203223213-1321320321320030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131233302222331-3320313223030031-3230132112013302-2000210103023301-3132000223221033-3011110330302002-0202131000221313-2320020300301231"></a>

## ingress_egress_gw.hub.spoke_vnets — spoke_vnets / 201133321132 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- ingress_egress_gw.hub.spoke_vnets

<a id="canonical-1000100130233331-1212021211112230-0200221230030332-0122032231003231-1001320013300020-2231203030002103-2303321232110123-2122133031111233"></a>

Type: `"list"`. Computed.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1331333002001000-2223020120320011-0132221030103203-0000011123013330-0113221201013220-1113003122201320-3313201113113101-3033133222031000"></a>

## Direct properties — spoke_vnets / 201133321132 / 3

- [auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-0010110221023231-2312230320010020-1230130300023313-3100022023321213-0233232231132112-1300010320222003-2020211300110220-3230200233323300): complete subsection reference.

- [labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-3302131032333213-1003222320312102-1312020222020030-1233130331332231-0232222201011312-2123200331113023-1010231113001100-1210003220230312): complete subsection reference.

- [manual](data-sources--azure_vnet_site--reference--group-004.md#canonical-3223120320231102-2313000302220333-1111133021211232-1220320310001321-1122022222322331-3310031333111303-2110011330212223-0203323331120021): complete subsection reference.

- [vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-2100023302303201-0021000303003231-1120222120002313-0102122032012101-0221121102302230-3100332131303220-2232032003112033-3000212032212021): complete subsection reference.

<a id="canonical-1022212001003232-2302301313233030-1211100221110112-0011131201221103-2012110022223321-0020222212123322-3123220302110320-2200302310201301"></a>

## Next pages — spoke_vnets / 201133321132 / 4

- [ingress_egress_gw.hub.spoke_vnets.auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-0010110221023231-2312230320010020-1230130300023313-3100022023321213-0233232231132112-1300010320222003-2020211300110220-3230200233323300)
- [ingress_egress_gw.hub.spoke_vnets.labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-3302131032333213-1003222320312102-1312020222020030-1233130331332231-0232222201011312-2123200331113023-1010231113001100-1210003220230312)
- [ingress_egress_gw.hub.spoke_vnets.manual](data-sources--azure_vnet_site--reference--group-004.md#canonical-3223120320231102-2313000302220333-1111133021211232-1220320310001321-1122022222322331-3310031333111303-2110011330212223-0203323331120021)
- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-2100023302303201-0021000303003231-1120222120002313-0102122032012101-0221121102302230-3100332131303220-2232032003112033-3000212032212021)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0010110221023231-2312230320010020-1230130300023313-3100022023321213-0233232231132112-1300010320222003-2020211300110220-3230200233323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213312131033003-2301022303132130-2021331232122211-3332232131323202-2313302302100222-0323003100111123-1101021103122111-1321211303311203"></a>

## ingress_egress_gw.hub.spoke_vnets.auto — auto / 000200331100 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- ingress_egress_gw.hub.spoke_vnets.auto

<a id="canonical-2003012003033103-2101313300120320-1032231321121302-0030323233230130-0111231323332201-2203230311120221-0132020031001313-2121100010230011"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-1321002132330203-2310200222231322-2031101323202031-0231013301322133-3222302101223223-3212320310001330-2023013223221010-0121100333311123"></a>

## Direct properties — auto / 000200331100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301003121203333-3300213312133011-0312132220220011-2122031130212330-2322313112030120-2132301013022110-0002103133330130-0103001230302023"></a>

## Next pages — auto / 000200331100 / 4

- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3302131032333213-1003222320312102-1312020222020030-1233130331332231-0232222201011312-2123200331113023-1010231113001100-1210003220230312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111021332032331-3103311210303330-0332012300312001-3303021202331131-3003300230132221-2013322310001123-0031002031232221-1123131011211120"></a>

## ingress_egress_gw.hub.spoke_vnets.labels — labels / 321120300123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- ingress_egress_gw.hub.spoke_vnets.labels

<a id="canonical-0032233333303310-0311210002222232-1012032003220202-0131033003223302-1213032033333022-3211301200201201-2230001000232332-1313223203330232"></a>

Type: `"single"`. Computed.

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Upstream description:

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

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

<a id="canonical-2111110132020122-3321103033133100-2020321031030230-1211122211133023-1301333012132333-3221232311121123-0302031112302322-3132000313310321"></a>

## Direct properties — labels / 321120300123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021110300002202-1103023000311021-0100110101201013-1020032102330113-2020200310022322-0013222012301003-2113121321303201-2012230133232111"></a>

## Next pages — labels / 321120300123 / 4

- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3223120320231102-2313000302220333-1111133021211232-1220320310001321-1122022222322331-3310031333111303-2110011330212223-0203323331120021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110310110200021-1121022021131321-1010231312233002-1332331303113112-2213210021212001-1021313110331231-0033112000023323-0233012121011131"></a>

## ingress_egress_gw.hub.spoke_vnets.manual — manual / 010233021331 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- ingress_egress_gw.hub.spoke_vnets.manual

<a id="canonical-2031313000301100-2013321012121330-0031021333020233-3010111000130201-2011011320002313-2320211300020121-3001130131123031-2310023011132201"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-3302013022312003-0230003023222333-3012130032012231-3030230130303201-1333010022131112-0321201132013333-3312210310103330-1200033310233012"></a>

## Direct properties — manual / 010233021331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213100203233222-0230023032312322-1003031003130320-2111121233122102-0030221012113120-1122022201312333-3303032223030201-2231022231012011"></a>

## Next pages — manual / 010233021331 / 4

- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2100023302303201-0021000303003231-1120222120002313-0102122032012101-0221121102302230-3100332131303220-2232032003112033-3000212032212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221020113331102-2132322330132132-2230220302130213-0303231313332022-0202211223233122-0300031001101000-3032332221330023-0333122322010322"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet — vnet / 012101102102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- ingress_egress_gw.hub.spoke_vnets.vnet

<a id="canonical-0000331300113032-1132011011231323-0310221330201012-2313113211023311-2313030020013131-2033202302103132-1310203220232001-2131233101320112"></a>

Type: `"single"`. Computed.

Resource group and name of existing Azure VNet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-routing_type": "[\"f5_orchestrated_routing\",\"manual_routing\"]"
}
```

<a id="canonical-2023110023023212-3131111200023302-3211212032131332-0212110303201302-2233302230201011-1322021233023123-2230120303002123-3231300122001321"></a>

## Direct properties — vnet / 012101102102 / 3

- [f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-0100102313032323-2303013201000110-0331301311101331-1203113323232022-0300223331202000-0022131101301003-0320333303012222-0020113333213122): complete subsection reference.

- [manual_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-3131321102021332-3312202122033201-3103103310130020-1001312120330001-3222112110132200-1311331213210110-1213233130112322-2221233102201020): complete subsection reference.

<a id="canonical-2231001033123222-2332123002200333-2112213000003013-0313121202333012-2201323100100010-3123021010132302-2300300110220213-1030032012320012"></a>

<a id="canonical-1131221320111201-3022333012031303-2300102303022123-3003201010100122-0213233102103303-3212222230232023-3332210231132321-3311021103100200"></a>

## resource_group property — vnet / 012101102102 / 4

Type: `"string"`. Computed.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1111231123222113-3211201232112113-0033301202130213-1123032013111103-1132200311203100-1203322101323101-0300002103130230-0123001003100032"></a>

<a id="canonical-2311212231330113-2331321203112100-2213112330201123-1020022322001122-1100222302030010-1200122202123312-2121122201013123-3211321221330133"></a>

## vnet_name property — vnet / 012101102102 / 5

Type: `"string"`. Computed.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3131113331122231-0331300320301022-0113100010231221-2330332003002131-0022211202303200-1311020023230102-3123231330233112-2222313201120012"></a>

## Next pages — vnet / 012101102102 / 6

- [ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-0100102313032323-2303013201000110-0331301311101331-1203113323232022-0300223331202000-0022131101301003-0320333303012222-0020113333213122)
- [ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-3131321102021332-3312202122033201-3103103310130020-1001312120330001-3222112110132200-1311331213210110-1213233130112322-2221233102201020)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0100102313032323-2303013201000110-0331301311101331-1203113323232022-0300223331202000-0022131101301003-0320333303012222-0020113333213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032303121000323-3320330201231031-2010111101030221-0000213001030302-1023313121131220-1313320122212320-3000120131331210-1230202310233320"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing — f5_orchestrated_routing / 230313322131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-2100023302303201-0021000303003231-1120222120002313-0102122032012101-0221121102302230-3100332131303220-2232032003112033-3000212032212021)
- ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing

<a id="canonical-3310233032203013-1221311030022102-2000111302323032-3233031032313223-3022131302203020-3000303321033031-3311021013330312-0001122013311030"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-2233311221231012-1132031100032111-1212311331312002-2003110022222232-1210120310123132-2020102003210223-3313332210030331-3020332021133303"></a>

## Direct properties — f5_orchestrated_routing / 230313322131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202230002112012-1321320123223220-2322030013230002-1002320012023122-3201213013122023-2310300221013313-1110233232111121-2212221121233011"></a>

## Next pages — f5_orchestrated_routing / 230313322131 / 4

- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-2100023302303201-0021000303003231-1120222120002313-0102122032012101-0221121102302230-3100332131303220-2232032003112033-3000212032212021)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3131321102021332-3312202122033201-3103103310130020-1001312120330001-3222112110132200-1311331213210110-1213233130112322-2221233102201020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210331333323320-0232332113130020-3312132022110102-3223113222331232-3231103012112122-0221000000001313-2232120333020302-1222131010203133"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing — manual_routing / 323233213313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100)
- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-2100023302303201-0021000303003231-1120222120002313-0102122032012101-0221121102302230-3100332131303220-2232032003112033-3000212032212021)
- ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing

<a id="canonical-0302221033213002-3330313231033312-0112213210202033-0033301021323110-3100332213112220-2120003133222211-3323222232323112-3132111112312131"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-0321223312323120-2013023003113033-0203313312031321-1002303101000012-0112213223002022-0321003303030133-3202213320330331-0333202120312213"></a>

## Direct properties — manual_routing / 323233213313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330113133012020-2020222331231221-2110030032100212-1021332200021110-2131220230130331-1301230113112302-3322332211233101-1020233203313022"></a>

## Next pages — manual_routing / 323233213313 / 4

- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-2100023302303201-0021000303003231-1120222120002313-0102122032012101-0221121102302230-3100332131303220-2232032003112033-3000212032212021)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312001020301111-3002301101131123-1232213020211313-2123321123332122-1030100103030213-2112122033102321-1113213121120100-1023120122302012"></a>

## ingress_egress_gw.inside_static_routes — inside_static_routes / 131112110201 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.inside_static_routes

<a id="canonical-3100020101101132-0223213022002100-1211300203232111-1220021130212000-2330010312200131-1223020011102123-3000323121012231-3111222200301332"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

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

<a id="canonical-0110322001100303-2303223132313103-1201300100003322-1201002101013031-1122203323200101-0020133131333012-0003110222330303-0320111320312012"></a>

## Direct properties — inside_static_routes / 131112110201 / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333): complete subsection reference.

<a id="canonical-2110003222023333-0100012020012331-3033332033231333-1100300331000132-0110202120000313-0212003332310310-1130003212232001-1221223222003032"></a>

## Next pages — inside_static_routes / 131112110201 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331002322130032-2203231331202230-0031021102202231-3112011232333303-3011210001100321-1031002102022232-1130220002032113-3222033100202120"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — static_route_list / 303002200220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-3232001311001230-1031003310113002-1212331220111320-1011200133101000-3021132123302203-0232302322231223-2013101001011033-1330130003221001"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3112320111223323-2002030320323200-3201031113310033-2301212220000101-3031333332212310-1323130131013200-3331131231003310-2001233232100322"></a>

## Direct properties — static_route_list / 303002200220 / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200): complete subsection reference.

<a id="canonical-2320313312212020-0020313100023200-1101022231210001-2222033210233210-3122031333331132-0123100033133102-1030030222121022-0211201130321222"></a>

<a id="canonical-0130112122231202-1102122033301311-2023230302321021-2130113302132221-0113333333010000-0312123222203000-0210201333122131-3313230110201203"></a>

## simple_static_route property — static_route_list / 303002200220 / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3113300030123033-2210313132202301-2102020000202212-1201113132122131-3012213320112301-2200311133110223-2100202222020333-3013310331232200"></a>

## Next pages — static_route_list / 303002200220 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033012033211033-3302130012332220-1300113131123030-2123231220002322-1023332021333010-1001210120032100-2222003011211322-2112023131000121"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — custom_static_route / 311212120211 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-0203210100231030-0310302110321011-0212330321223132-0011130000333131-0121311032013202-1320013202103322-2310131333223131-2102000230231313"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

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

<a id="canonical-3011213121133012-0201232030031020-3102333311110011-3033101031001033-3220113230120100-1012102232010112-1221020112310023-2201001221010111"></a>

## Direct properties — custom_static_route / 311212120211 / 3

<a id="canonical-3103210132113320-1300211210011133-2012220003001333-0233313312031113-0332033302310210-2101013322110003-3331221302012103-1003312022303132"></a>

<a id="canonical-0323201012011121-2222222133312200-3103233302233032-1202133032012333-3232310111321122-0312120013010310-1210332113223301-0122121330222300"></a>

## attrs property — custom_static_route / 311212120211 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-2313202230023112-0020113132003330-3313332230120323-1330322322002102-2330023320313113-1230031013012332-3322131110330020-2031223013203021): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-0310203132231221-3321301322313320-0302010010320022-3312030213321011-0023313322121311-3303320200111322-0201332101232023-3033213202300320): complete subsection reference.

<a id="canonical-0013021333211133-2101100330111021-2123112031121210-2222323032113233-2010032221122223-0123202123133132-1330322022322230-1122013020020101"></a>

## Next pages — custom_static_route / 311212120211 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-2313202230023112-0020113132003330-3313332230120323-1330322322002102-2330023320313113-1230031013012332-3322131110330020-2031223013203021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-0310203132231221-3321301322313320-0302010010320022-3312030213321011-0023313322121311-3303320200111322-0201332101232023-3033213202300320)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2313202230023112-0020113132003330-3313332230120323-1330322322002102-2330023320313113-1230031013012332-3322131110330020-2031223013203021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310033010332322-1212000102302033-3000333032212300-3022221213022233-2012102000133301-0211230222210333-1102232332311013-0032003331100100"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — labels / 120220020332 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-0301202033123320-2021222122320122-3113111020112312-0301032131021300-3222302132321023-0112210201210230-0120310201203100-0110031211331033"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

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

<a id="canonical-0113320311130101-2011021330301211-1323200021213210-0312320100003012-3012313211000001-2212133111200211-1100323103120130-0312302031203201"></a>

## Direct properties — labels / 120220020332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022232210103213-3333032213312213-3330000012101013-2232111131030023-0313303233010020-2002210322201003-0213002230003030-2303030330203233"></a>

## Next pages — labels / 120220020332 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320122023033213-1231200233020312-1302202122120021-3330223212302321-0210233232100220-3001230130013232-3223220000103111-1011323021311320"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 301231301220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-1021302330130233-2030011023132313-2020003002013321-0100222200321031-3113333232132233-0330232300101032-3012313330330331-3200122123022030"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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

<a id="canonical-0231033030311213-3033302301313302-0211000131121312-3321121303130200-2133222221002223-3022133121103111-2202302212001121-0133123301311011"></a>

## Direct properties — nexthop / 301231301220 / 3

- [interface](data-sources--azure_vnet_site--reference--group-004.md#canonical-3302223013303103-3202102023221102-2010211011320213-1130211001211113-0310002301320301-1231130132232303-3030120020201201-1101323020012110): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222): complete subsection reference.

<a id="canonical-0323001030201201-1302103310310332-3133223211222022-2300303201100213-3010000333030230-3111030333133201-0220303102111231-3103120211330122"></a>

<a id="canonical-0121312031301100-3200302221332031-3020103133231230-3323301023202221-0011032321200201-3231313131113303-2320011233023222-1213222002321110"></a>

## type property — nexthop / 301231301220 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0132322103311320-2313200023332101-2010110023300200-2110200001201311-1130333213022322-1011331100222323-2020221302012321-1132230311321100"></a>

## Next pages — nexthop / 301231301220 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-004.md#canonical-3302223013303103-3202102023221102-2010211011320213-1130211001211113-0310002301320301-1231130132232303-3030120020201201-1101323020012110)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3302223013303103-3202102023221102-2010211011320213-1130211001211113-0310002301320301-1231130132232303-3030120020201201-1101323020012110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332333313310011-0311211020023322-2232333020122002-1333131122000120-2331322100302031-2330201330031110-2033301321232010-2202302113133230"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 213131122301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3002331122330310-1322322010232123-3220012130032110-0122232013232233-3232232021311320-2223031111032221-2222131123010312-2130133212020231"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2100130203122313-2200010233111110-1211112110121330-3300120111010202-2233100121101232-1200123300012003-0230130331111233-1322310003323130"></a>

## Direct properties — interface / 213131122301 / 3

<a id="canonical-1030312230321220-3221320210223313-2102320232032100-2030023132130133-3121312201303332-1030203103020012-0113003121020301-2233032001032101"></a>

<a id="canonical-1201331012312333-3121103003132033-1121130213230201-2331211313030213-2312331311003001-2032333313112021-1110011011201003-3222022321233131"></a>

## kind property — interface / 213131122301 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1332121310033221-1320102133130200-3123323311001111-3130322030120303-3132333331210022-2121210233321132-1113310002301321-1200313102200012"></a>

<a id="canonical-1301202202100111-3212320222130210-1211210313203332-3033221223200000-1102200120312301-0110321010213331-2131102000313012-3120220001121313"></a>

## name property — interface / 213131122301 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223010112220100-0030230213333020-2313331033131032-3013313132123121-0223123103221122-3211022303210312-2201022122120332-2102123123220110"></a>

<a id="canonical-0210130220210012-1321023302020133-0201200311202012-2031332223201033-3221103213003210-2330101123010103-3333323311130220-3320123103121202"></a>

## namespace property — interface / 213131122301 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1003103201122310-0223330013103302-0011313020120203-1032230223131313-2111001232323023-2232222130202103-3131110021301100-1311303213200232"></a>

<a id="canonical-1032022211323023-1011011323200200-1121121213112013-1320331231330130-0331211121132010-1331110110231102-1312120010002333-3012200012213331"></a>

## tenant property — interface / 213131122301 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0202012221230032-2020020030231302-3102021032000311-2210023111213320-3331203201002120-3103013223220310-2102212001313100-3313230000203001"></a>

<a id="canonical-3301203221310302-3332233133103110-0202122102131322-0222133201002210-2101320330211001-1110331121030000-1302301220011020-3011022000112223"></a>

## uid property — interface / 213131122301 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022103100323331-0333032320003023-2210011202303201-1303213302313103-3322103130131200-0210011303212100-3311121213230231-1000321332303222"></a>

## Next pages — interface / 213131122301 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302310200112321-1032021302320203-2232202311120211-1020313222232031-0201032033313103-1313303133223120-3300033111332222-0032120320102202"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 303322010330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-3012021323002302-2101332101013103-0231132320111212-2211131210130103-2313322101200132-3222200223113123-3022323203130311-1000213113201211"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-2033312022322130-0012320311121032-1110123232310233-3002011100223112-0303121122120112-3110131202302133-3322110100020332-3012130332102130"></a>

## Direct properties — nexthop_address / 303322010330 / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323): complete subsection reference.

- [IPv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-1020103122133232-3112313133303033-2222110110032331-1103001000020100-1321022011002110-1202011000203023-3101221131231031-1103302013212300): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-3112201300321130-0110132022022002-1200133221203022-3113213012023322-3202111331221200-2213211031203100-1101211200000021-0011320122233120): complete subsection reference.

<a id="canonical-3212230331223332-3221002031102332-2303303020232012-2202031113332032-2331112331013102-3003310233311302-1113230021110112-3122322030322311"></a>

## Next pages — nexthop_address / 303322010330 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-1020103122133232-3112313133303033-2222110110032331-1103001000020100-1321022011002110-1202011000203023-3101221131231031-1103302013212300)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-3112201300321130-0110132022022002-1200133221203022-3113213012023322-3202111331221200-2213211031203100-1101211200000021-0011320122233120)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321102121220233-0000110013232232-1032123120321333-3302132300021001-1302231200123021-1120102132322113-3232200103102122-1120120101130013"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 231121113133 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-0320132013322220-1332231030303300-0133223010022213-3212103312020321-3300210302032111-2112123302020310-0110200010210130-1011332100021112"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

<a id="canonical-3003120301100103-2021013301320121-0210322010213132-2322000232220333-2310313310321121-1132322000131223-0111230230302320-2000011210221113"></a>

## Direct properties — dual_stack / 231121113133 / 3

- [IPv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-1323300110220112-3231022333133233-1132322212231111-1011200201030313-2232030312212031-3121212102232313-2323022032010131-2203002310323321): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-1033331123222123-3302012111222332-2013302123210303-2203322321001120-3312221330331313-1232211333231223-2011200320011232-1020210120200311): complete subsection reference.

<a id="canonical-1300121312202300-3122211312332003-0122200022111011-0122012131211111-0332223320023323-3102103322332003-0211321302232113-2132121122030113"></a>

## Next pages — dual_stack / 231121113133 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-1323300110220112-3231022333133233-1132322212231111-1011200201030313-2232030312212031-3121212102232313-2323022032010131-2203002310323321)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-1033331123222123-3302012111222332-2013302123210303-2203322321001120-3312221330331313-1232211333231223-2011200320011232-1020210120200311)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1323300110220112-3231022333133233-1132322212231111-1011200201030313-2232030312212031-3121212102232313-2323022032010131-2203002310323321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301033323211123-3232332131102321-1233120023222213-1130221111230211-2110330000111100-3011211122121200-1223213223221222-3212012010123313"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 222311320200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-0122330210321112-2223221131333332-3332033210330010-2332321322300013-3311000330112201-0323320302023033-0023323113130032-2331302031312200"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-1213211113002102-2102321131030121-2330211303000230-1203012221112112-2312112012302311-3312332332222121-2232131132331322-3223022222011323"></a>

## Direct properties — IPv4 / 222311320200 / 3

<a id="canonical-2322302100321122-1013020023211131-2231321100301230-1132013300333002-3120132112113013-1002133031221200-1102220202110232-0023033313123001"></a>

<a id="canonical-0031200111232102-0122123320213323-1001101103110300-0202210313301322-2301101201112031-3122023313012003-2310012100122211-2120103212223312"></a>

## addr property — IPv4 / 222311320200 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3332320233002030-2003230332300020-1333023110013201-1303223223013231-0301011032210123-3003132101122002-2022313112223203-3221332012103122"></a>

## Next pages — IPv4 / 222311320200 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1033331123222123-3302012111222332-2013302123210303-2203322321001120-3312221330331313-1232211333231223-2011200320011232-1020210120200311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222101230233232-0220122111100203-1030232330121012-0211223313002023-1102131031210312-3101023031210032-3210131311331210-1000031212123331"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 321203212130 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-0211020201111310-2031313032213011-3003031100232123-0321211130323113-2001331331133323-1203131020320010-0321102010321112-3130020232023032"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-2101022202233032-3333303221133131-2022120210303331-1000002322322302-1220131202000212-0212131120203102-3110113220001200-3202020110213223"></a>

## Direct properties — IPv6 / 321203212130 / 3

<a id="canonical-1331120013232330-1301220223322121-2012121032212012-0232212233021230-2330020112210003-2011211120321313-0001013113331303-2032213022313300"></a>

<a id="canonical-0013120113010300-1113112311022301-3213203023133023-2101110233313230-2230013232300211-2310210023332323-2332100023113300-1113230302033113"></a>

## addr property — IPv6 / 321203212130 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0110213233020222-2330223012130333-0031021201131030-3232003223112321-3230003330313332-2322320213331031-3021023200110022-1200003212200320"></a>

## Next pages — IPv6 / 321203212130 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1020103122133232-3112313133303033-2222110110032331-1103001000020100-1321022011002110-1202011000203023-3101221131231031-1103302013212300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212102123323232-0012012013320331-0102032010303212-2120210030111013-0021100022233020-3003102103012321-1100102300110010-2232232132101023"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 131210211301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-3211202132311313-0033200311101002-2201231213212320-3000132020223002-2021333112120211-1323232011102023-1201112000113211-2021103130222300"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-1000233232201313-0013001110211110-2033322322013003-2131233312321103-3312112201131111-2111101010123133-2331331230303201-0220300202103300"></a>

## Direct properties — IPv4 / 131210211301 / 3

<a id="canonical-3231110132312131-0101330112030031-2102203312203233-0312021003031220-2123012310303132-2123210113331330-0231103101020022-2010202233303111"></a>

<a id="canonical-2213212213320002-1020101020122112-2001330031213121-0031120300001322-0122023130111002-0313010303032002-1001120113230100-1013200312012103"></a>

## addr property — IPv4 / 131210211301 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1201201122230231-0231133213132010-0233323113313132-2303002022331033-2322131221103112-1310210003323231-0113100102331011-3230313103111213"></a>

## Next pages — IPv4 / 131210211301 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3112201300321130-0110132022022002-1200133221203022-3113213012023322-3202111331221200-2213211031203100-1101211200000021-0011320122233120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310332323322320-2010101112323023-3133013210131323-0111222011311102-3112322122012300-3132102112301121-0312223231033222-2323031101231202"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 113022331123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-2032330001300233-2003213001211202-2013121012313231-3021021333321130-1300021302221003-2232301203111210-2300113330131111-1100312120332033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-2003331320231132-1333312331310101-2312023313132021-2113012203210120-3113233310203113-2320322331131011-2132210013101030-0003022332112322"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-1023012210100012-3102211033311020-1213230013200122-2311031230323230-0232321213121010-0200001213020213-0313322033131113-0130232311331210"></a>

## Direct properties — IPv6 / 113022331123 / 3

<a id="canonical-2333002110012002-1001312110331203-3113023033311210-3002232220212200-3230301001233030-2310032303010310-3020033332120120-0232033100130200"></a>

<a id="canonical-3031202231233313-2013001231030232-0233122111103033-3232200211212322-0022111332122220-3230113123022112-3133222332212333-0013203131321201"></a>

## addr property — IPv6 / 113022331123 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1320230301102001-0200303222210120-2230220232121013-1302033122210210-3133123222322301-1211100012203330-0301311202230313-2333322021122213"></a>

## Next pages — IPv6 / 113022331123 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-3130003310321311-3302310300220012-0333023011301310-1220011301200132-3311000200322333-1010210013301131-2222113001113030-2112200303323222)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0310203132231221-3321301322313320-0302010010320022-3312030213321011-0023313322121311-3303320200111322-0201332101232023-3033213202300320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131313132220331-1132202231220100-3130312302113002-3122320112131321-2332102030012133-1230132002220030-0033013033211110-2331211212333021"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — subnets / 203122233202 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-2003130211113302-1100230320122230-0332112021220331-0212312311311333-2032333002320332-0230010313330212-2321233213313133-2311010232332021"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-1021320132021302-2111231131313331-1300210132003021-1110200221013310-0032001233323310-1331212111323012-2020133330310332-1312213202123101"></a>

## Direct properties — subnets / 203122233202 / 3

- [IPv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-1233010311203301-0311210101003220-2131012320032001-3112023300333103-1033101121001030-0303301033131303-0023202313232321-1120231023300132): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-1022031122032131-1221111210123120-0032131202111223-0303023200232230-3203133312111312-1202130113132300-2030321232132320-3032023212301313): complete subsection reference.

<a id="canonical-3020131223232321-0200322220301311-0222301210302021-3102100231133003-3012322131321013-2013013230212313-0112011023313231-0323001012231202"></a>

## Next pages — subnets / 203122233202 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-1233010311203301-0311210101003220-2131012320032001-3112023300333103-1033101121001030-0303301033131303-0023202313232321-1120231023300132)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-1022031122032131-1221111210123120-0032131202111223-0303023200232230-3203133312111312-1202130113132300-2030321232132320-3032023212301313)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1233010311203301-0311210101003220-2131012320032001-3112023300333103-1033101121001030-0303301033131303-0023202313232321-1120231023300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311330010013221-2230333023003212-0322203122210111-2022131202013201-3121020110332003-0031110231220133-2131323001212100-3131103222223032"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 310021221312 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-0310203132231221-3321301322313320-0302010010320022-3312030213321011-0023313322121311-3303320200111322-0201332101232023-3033213202300320)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-3003230033030122-2313302010211320-2312203321313330-1321312000120221-2122300231033001-3001030001133101-3332222302221002-2113031121221113"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

<a id="canonical-2130120210210203-3301231301032123-2323300101111023-1003232203233031-2022001102020300-3322103212030033-3030003102212221-3030112013032222"></a>

## Direct properties — IPv4 / 310021221312 / 3

<a id="canonical-2312002102113013-0223122232220303-0103010330122020-0110222130312230-1233020330102223-1320311301111123-2011333211231012-0211130223002220"></a>

<a id="canonical-2010323210010010-0213012211211310-3211012130323221-1110101223200033-3122000133223331-0201133211130130-3000130020022302-0220122333010223"></a>

## plen property — IPv4 / 310021221312 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0233221200311232-2020330120202330-0231322132313122-1101213212123311-2333022130212220-3231131321310001-1031120311331013-1203013131032100"></a>

<a id="canonical-3200100022003330-3102211200030230-3020012121331333-2322130133031120-0030132211102010-2101210013020011-0213311331221332-1033000311031010"></a>

## prefix property — IPv4 / 310021221312 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2202320223211233-3131131013133033-2113011121012310-2301223320333030-3032313113130122-0220221212333223-3222221023002133-2200130013030030"></a>

## Next pages — IPv4 / 310021221312 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-0310203132231221-3321301322313320-0302010010320022-3312030213321011-0023313322121311-3303320200111322-0201332101232023-3033213202300320)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1022031122032131-1221111210123120-0032131202111223-0303023200232230-3203133312111312-1202130113132300-2030321232132320-3032023212301313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002031113133332-1133332201121130-2010000321331300-0102333010320331-0022303001030322-0120313232301110-2102033103120312-3131102332202300"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 223310231303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-2230213201103212-0210220122233211-1102313002130100-1001230300322303-1232132321301002-0132221331123221-1203030203213323-2001002303321131)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-0212003020321232-3021211122330331-2323221022111111-0333222223023103-2330220230101201-3101130212213211-3320000030213001-2301111133322333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-2132211321323100-2021310232111021-3303211313132210-3120230320023012-0012222022003012-2331101301211211-0103310310102202-3201220021102200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-0310203132231221-3321301322313320-0302010010320022-3312030213321011-0023313322121311-3303320200111322-0201332101232023-3033213202300320)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-2023323012012032-0330110023322023-0031111201221113-0221222232200010-1322123001321101-1110211110330322-2013123303130333-1233200031303031"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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

<a id="canonical-1220011012022023-3131312332113023-2231231321021212-1123212330101030-1221330232321333-3000323313310000-2111300001032131-3301012030233220"></a>

## Direct properties — IPv6 / 223310231303 / 3

<a id="canonical-2231301302332132-0030131202313120-2332211220321012-2200033002213200-2302221000101130-1100300001003213-3210213121110002-0101231201333102"></a>

<a id="canonical-0200000310122211-2213032120013020-1032201330023222-3220323202000103-1222221212323101-2032000332010222-2322030213222010-3003313212222203"></a>

## plen property — IPv6 / 223310231303 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1210323312321123-0023212001123300-3202233123010212-0222023012231102-2310002113331320-0120213022203002-0323120222302333-3221232002222131"></a>

<a id="canonical-1122133312012203-2203303311012231-2110131331121301-3122310323102132-2000210320112233-1000303003001221-1001201022203320-0231030131220203"></a>

## prefix property — IPv6 / 223310231303 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2000301023130021-3221011213130333-2232023222212203-3101301003203111-3112100022033232-1201213301022300-0012122320202313-2212311202203230"></a>

## Next pages — IPv6 / 223310231303 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-0310203132231221-3321301322313320-0302010010320022-3312030213321011-0023313322121311-3303320200111322-0201332101232023-3033213202300320)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0101111323330133-3003201021331100-0310122320232331-3203001323210032-3320321202122102-1021233330323013-0323110131020321-3103130221120020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033020213301212-2333033101003113-3033322132130123-1231331113332223-3032310203020032-2021212030013031-1113120312010132-0130213323033103"></a>

## ingress_egress_gw.no_dc_cluster_group — no_dc_cluster_group / 203223332301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-3231111223103113-1022231333033211-1012123011120213-2002301031320320-2331123332323302-2213231032003212-3202121013331211-2003001322001120"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-0121030133122303-0313302120013111-0231321110033300-1102212133133132-1331202303231031-2321203231333230-1310011233121131-3212213230133301"></a>

## Direct properties — no_dc_cluster_group / 203223332301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233023131323121-0020220210020031-1100212223130200-1121003131132023-0310031122031302-2120222122330333-2220031320303202-3301301330121102"></a>

## Next pages — no_dc_cluster_group / 203223332301 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3103031102110202-2021210100021303-1001211330200010-3233032000130200-3203301313113230-3031030113200323-0212121120100102-1231100323203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121211121001022-3201320300103100-0103130032212213-0311023313323303-2212023212021220-3110332212110013-1321303302001313-2110011300202210"></a>

## ingress_egress_gw.no_forward_proxy — no_forward_proxy / 202021200210 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-3122200311333211-1213312222012113-2222113020000021-0001212133320122-0220130012102222-0010112112323102-1200123021220221-3021110203002033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-2030102000201021-0012101023210301-3122020231313301-2021333130123220-0203032123303121-0203021022331013-3123201112032213-2301230020322131"></a>

## Direct properties — no_forward_proxy / 202021200210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210312303110201-2011323202330010-3112320011100313-2300303110300232-0212101223313011-1233102033321032-0133133023120102-1130223133100111"></a>

## Next pages — no_forward_proxy / 202021200210 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2303123333010200-3311230023001322-3230033332122221-2232333220333300-1101002221130011-3131031023303121-0031232031200133-3301200331023320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113023202120111-1231311121201300-0200111203023233-1330231123033101-3130330022120131-0211332321221131-3103101122222130-3122223120233203"></a>

## ingress_egress_gw.no_global_network — no_global_network / 311331101113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.no_global_network

<a id="canonical-3013103300212233-0020233023110230-0202100220223230-3313201200320110-0300213232220202-0301011233020333-2101130302100202-1301211201303322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-1120222222101232-2033311030122121-0320022120322132-2211321103332003-0202220322323131-1313030300333320-0202331000301230-0031300310321323"></a>

## Direct properties — no_global_network / 311331101113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223322020332201-1000032320300330-0123202333303303-1203311212102121-1311102101330002-0312020032103202-0313032011232011-2222333002322210"></a>

## Next pages — no_global_network / 311331101113 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2310020312010031-3210033123023322-3123120332231303-3200323123322300-0131031021031030-3120213130023010-1132010131210033-0313030313122022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221021303120130-3203201323211001-3200111311223230-0032232103331131-0101130123320110-0011102001330311-2322223312322331-1200013221121210"></a>

## ingress_egress_gw.no_inside_static_routes — no_inside_static_routes / 322212202032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-0002212022111021-3202321033201022-1220033122021030-0103200233323321-3002121231211010-0032311021221313-3121100303231021-2032020211123210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no inside static routes.

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

<a id="canonical-1220230312211333-0211310232312101-1100022232031202-2211332202021312-2030032031203221-0003303103030002-3223222011232033-1202013030232030"></a>

## Direct properties — no_inside_static_routes / 322212202032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220100310313022-3223123221322331-3310123302130003-3331010211232313-1011033330232123-2100223202012322-0213022123110132-1011203333213302"></a>

## Next pages — no_inside_static_routes / 322212202032 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0322132003301000-1221211302010200-3202011220201211-0032101001220010-3201321332133001-1020132313001310-1333103300322112-0203012301301233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221102003222301-0330023112212000-2201001323233330-1302001003232001-1023213031132211-0130033330000020-3333113223203311-2010011322221320"></a>

## ingress_egress_gw.no_network_policy — no_network_policy / 103112313110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.no_network_policy

<a id="canonical-0031313201001322-0301220020132331-1020222313230212-3332210203132032-3010011320111011-3032333233002201-2133130031333102-0030231210211213"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-3311213100332233-0011032233321230-3101311200101110-2221310002322211-3131011131013231-0011112330230100-1303232310330232-3301210311120112"></a>

## Direct properties — no_network_policy / 103112313110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320123313210213-3302001013232303-2033111132003010-2202011113313101-0011113133323221-0203210321021000-0131220001302020-2303203011130223"></a>

## Next pages — no_network_policy / 103112313110 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0000022213123000-1123031211123111-0232002100210130-2102133102110313-0211010130022331-2320221203210030-2223333012220013-1123002123212102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222230132113021-0221221113210203-3123201101132202-2032012013112230-3333012220211331-1122311201300031-3232201212313003-2131323311002133"></a>

## ingress_egress_gw.no_outside_static_routes — no_outside_static_routes / 032223201302 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-2112010010021232-0112201001003303-2203110332001310-1233322223112300-3222133301001003-3231121223130032-2030100223100112-1022021213303330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no outside static routes.

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

<a id="canonical-1201333021311101-1302002023032022-3112210221331102-1102322031022311-3302113010023113-0033103220222033-1023203233330032-1010212210330032"></a>

## Direct properties — no_outside_static_routes / 032223201302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131200122002110-0131112100233012-2030111210113012-0231003333312001-3320202012332223-1100220221000110-0100302031223303-3233203030230012"></a>

## Next pages — no_outside_static_routes / 032223201302 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3203233020233221-3322303233211320-3223321223301033-0103001230003132-3233020201011301-3312331121030113-2030101121330012-2120030120120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031330032022010-0210000020001211-1301020233231100-3110030103333133-3311222333231131-0010301203110300-2102320112213132-2123021331220112"></a>

## ingress_egress_gw.not_hub — not_hub / 001212111112 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.not_hub

<a id="canonical-0323132003110033-2221111011121002-2201322331022212-2022202233021113-3332311210130200-1112230223313211-3011223331212021-0003022323103321"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-3330130133111311-0313132233113002-0320203303321123-0022031303321310-0102012230101110-3301131312003123-1122002113100231-2132011213023012"></a>

## Direct properties — not_hub / 001212111112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030232230303003-3233210100213011-2323120231031010-0031020001113312-1231203013302211-1100031111020313-3002230211031003-3311121202000133"></a>

## Next pages — not_hub / 001212111112 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311231330110122-2222010021332033-0001020323302131-3200132121233120-1202003333230011-3200101012330112-3102032330113110-0232201102323112"></a>

## ingress_egress_gw.outside_static_routes — outside_static_routes / 333211021310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.outside_static_routes

<a id="canonical-2130332000033112-2222001322001123-1202003113001211-2201200200230113-0320113121311230-3222313203330101-3220131102112311-3220001120103313"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

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

<a id="canonical-3312303012310321-1311322302220031-3131303220003333-2231302220102110-2023222211220323-3321222111200202-2012233101312311-3303023210220100"></a>

## Direct properties — outside_static_routes / 333211021310 / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311): complete subsection reference.

<a id="canonical-1122301011211111-3121001232121112-0111101111100231-2221321231000131-2123333122220100-2203212010232321-2332222323333313-1000112002311131"></a>

## Next pages — outside_static_routes / 333211021310 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310211033123121-2223121013002332-1122112311133223-1330330322201113-1311111222120331-3001120103331333-2233001233110223-1100233221000131"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — static_route_list / 302033300111 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-2123112030000323-0303201323030002-0132200002120001-1033233113201012-2223322100010221-3103200211010302-0130003333110032-1230122203012332"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1303102310110320-2213101323031221-0112331010122122-0222221331321311-2003223013230121-3033201120012003-0122010222301320-1321033032233113"></a>

## Direct properties — static_route_list / 302033300111 / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300): complete subsection reference.

<a id="canonical-0331100212130122-0333300203030322-3133320103103123-3100222322001221-3121120101013322-1221010220011303-3333011312010132-0023100310010210"></a>

<a id="canonical-1011133011222001-3131232311111113-0312231311110003-0330213011000312-0133001021110220-3222113112221030-3230231312101301-3301313223203211"></a>

## simple_static_route property — static_route_list / 302033300111 / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```
