---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-2021013222121011-2320121131001021-2010000101200222-1103323230323310-1023202232311322-1023321231211103-3133321002210022-2002003133021112"></a>

## Next pages — values / 033212300133 / 8

- [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-3003321210210000-1031301202332201-2131021313332210-1112031010202333-1321123133123000-2330321312312210-0133321220130230-1113102030120101)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2032330000021121-3311023021112132-1120113313102232-3121212313333300-1010212230003122-0321211321302330-0320223133020322-2302030220001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111232211111121-0321131310033311-3200311330311230-1310213332333120-2110313301330001-1200130301110222-1122212002200101-0131133212213201"></a>

## primary.default_rr_set_group.cname_record — cname_record / 131320102031 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.cname_record

<a id="canonical-3213013112112213-2233222323123300-1300131223222322-3231230120300333-1003233213112103-1020331020111311-0120221020321123-3203313311231200"></a>

Type: `"object"`. single nested block, Optional.

DNSCNAMEResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
cname_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233013121320333-3221110331310212-1233223123131303-1120030301233020-1212123200232012-3220310201101200-3233002022312203-0000131223120101"></a>

## Direct properties — cname_record / 131320102031 / 3

<a id="canonical-0133022313203033-2203213132101313-2012320213100012-3132220123330222-2000312122221010-0120121301001232-0112112212302123-2202213310202020"></a>

<a id="canonical-2033003102311331-1012121112302303-0002011323003121-2121231212233010-1010320120230220-3103010333203102-0001310023002321-2101312100133011"></a>

## name property — cname_record / 131320102031 / 4

Type: `"string"`. Optional.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0031110333131103-0321200112203030-1131333320211332-3100303112030300-0313332122231113-0203112203112102-0223101233133010-3021121003102131"></a>

<a id="canonical-0120331202133201-2003100012001312-2211222021322311-2310101013303223-0330032302203022-2131200211312332-3301000322213031-3213012022313111"></a>

## value property — cname_record / 131320102031 / 5

Type: `"string"`. Optional.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-3102331120102221-1311000120300203-1220013023121330-3132120322230300-2103332312310230-2303322212333111-3032002330100201-1003303320111113"></a>

## Next pages — cname_record / 131320102031 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320310030212232-0120233013000102-3113112323111221-1020001012113110-3032232111110123-1130220212010202-3312332331231120-1123123002131222"></a>

## primary.default_rr_set_group.ds_record — ds_record / 221100023100 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.ds_record

<a id="canonical-2122200203213103-3312233120033010-1121300202323310-2012331030031320-0202011320020000-1031332320320322-1131032103120030-2100310212112121"></a>

Type: `"object"`. single nested block, Optional.

DNS DS Record. DNS DS Record.

Upstream description:

DNS DS Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
ds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201213213210203-0311230323021101-2121210230113212-1323210332020232-3031311111200300-3320321322031322-2011312312300221-0312200231022321"></a>

## Direct properties — ds_record / 221100023100 / 3

<a id="canonical-0003003332132021-1232010212023013-1202202122212203-3321002101310211-2032030102010333-2332022330123010-0100200233021103-2101112313313201"></a>

<a id="canonical-0110030201322100-2202031123311123-3312100103230033-0331310103003023-2021312200102112-2121323330111121-2100130203111110-2333001213132212"></a>

## name property — ds_record / 221100023100 / 4

Type: `"string"`. Optional.

DS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312): complete subsection reference.

<a id="canonical-0200321010001032-1022011312113122-3002012320321232-1132022111231021-3003111022132020-1103012212311321-0013112132300223-2103230310321030"></a>

## Next pages — ds_record / 221100023100 / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232203302001111-3022213300023212-2101202023012302-1023130110300010-2121303231123333-2323031323200013-0233112120110102-1200130033103101"></a>

## primary.default_rr_set_group.ds_record.values — values / 211102201211 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303)
- primary.default_rr_set_group.ds_record.values

<a id="canonical-3203301320223220-3031220203300100-3132031200201333-2121030223322223-0012323033230132-2010022022203233-3000201203121101-0300300331031312"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301313332302003-2323023221200100-3313322100232121-0221113013031301-1123332131200110-2331313112033112-2221223132322321-1312023032102233"></a>

## Direct properties — values / 211102201211 / 3

<a id="canonical-3130200110322321-1033133210111000-3121132222003312-1310323332002003-3120112332202223-0320231203013310-0211302003333320-1312212222002301"></a>

<a id="canonical-3003012112213131-3010023002103110-3011200002001033-2023331223003222-2121311301212230-0012331212320323-1133113112311320-1011303121231313"></a>

## ds_key_algorithm property — values / 211102201211 / 4

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2300310211211003-1123233032120003-3023032130011130-1301020002001301-0101122302120000-2113101202322213-1012000203212322-2213132110133321"></a>

<a id="canonical-2120033132312130-0311310312011211-0201223313322311-3032031212330030-2030101301100301-3220022030003230-1103233131102302-1023313120320010"></a>

## key_tag property — values / 211102201211 / 5

Type: `"number"`. Optional.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](resources--dns_zone--reference--group-002.md#canonical-3330103022123210-1130322031123213-3231302220122211-3201121100311001-2130012321201132-1022313330023003-1203222322231021-3323212301010001): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-002.md#canonical-3213221210303123-0123212131003113-1122132010223003-1313212133221102-2002211213210333-0012022112211310-2333311233332200-1301322122122122): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-002.md#canonical-2210121320301100-3213010231111311-0223132131103032-1030023212322202-1222200221131333-3100002232003002-3210211231131031-3123013311020031): complete subsection reference.

<a id="canonical-3010332230111112-0032021330123212-3032012111020312-3032220331113111-2121122110011003-1030112113203321-3021323201113030-3023210111020020"></a>

## Next pages — values / 211102201211 / 6

- [primary.default_rr_set_group.ds_record.values.sha1_digest](resources--dns_zone--reference--group-002.md#canonical-3330103022123210-1130322031123213-3231302220122211-3201121100311001-2130012321201132-1022313330023003-1203222322231021-3323212301010001)
- [primary.default_rr_set_group.ds_record.values.sha256_digest](resources--dns_zone--reference--group-002.md#canonical-3213221210303123-0123212131003113-1122132010223003-1313212133221102-2002211213210333-0012022112211310-2333311233332200-1301322122122122)
- [primary.default_rr_set_group.ds_record.values.sha384_digest](resources--dns_zone--reference--group-002.md#canonical-2210121320301100-3213010231111311-0223132131103032-1030023212322202-1222200221131333-3100002232003002-3210211231131031-3123013311020031)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3330103022123210-1130322031123213-3231302220122211-3201121100311001-2130012321201132-1022313330023003-1203222322231021-3323212301010001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313013330332211-3310113312031221-2210120223100211-2223131330120012-3033032301302123-1330203231021312-3233310120331310-0001331222211312"></a>

## primary.default_rr_set_group.ds_record.values.sha1_digest — sha1_digest / 330232120123 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303)
- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312)
- primary.default_rr_set_group.ds_record.values.sha1_digest

<a id="canonical-1212003013102330-2302112011102223-2031011033332313-0120232331313110-1011332213222321-2213313320031000-1112130122030203-0001301012233012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211320103221230-1132333133013312-2333232023123220-3323110101321032-2220323312013322-2312221111003002-3201102110121232-3133203123222132"></a>

## Direct properties — sha1_digest / 330232120123 / 3

<a id="canonical-1133310021010122-1303322332310212-0331233303112333-1103321332321221-1003103332133312-3321102232021100-0010110123312033-1001322030132132"></a>

<a id="canonical-2030022300101111-3301011103101110-0123302213003102-3331212033133030-3132123110021220-0330103020331223-2122033031230023-0331013103010330"></a>

## digest property — sha1_digest / 330232120123 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-2221031013121300-2211320201333020-1100123120320331-0330101321211303-2120033023120101-2122310021033231-0201123331321312-2131032103223001"></a>

## Next pages — sha1_digest / 330232120123 / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3213221210303123-0123212131003113-1122132010223003-1313212133221102-2002211213210333-0012022112211310-2333311233332200-1301322122122122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320120010333131-0312220212321032-1211113213030303-3323223231121022-3310032130031133-0032002132210223-3213022113321101-0022113022302323"></a>

## primary.default_rr_set_group.ds_record.values.sha256_digest — sha256_digest / 221233030030 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303)
- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312)
- primary.default_rr_set_group.ds_record.values.sha256_digest

<a id="canonical-2321133331330211-2300110231313303-1213311223110312-2000323201331311-1022333011031103-2112113320300001-1103323000212033-2331220200320120"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321320320200212-0111031022021210-1212133023213323-1020023233122010-3202133103202223-0333220310202001-2103200322323222-3011023111222200"></a>

## Direct properties — sha256_digest / 221233030030 / 3

<a id="canonical-3031003211221223-3302222231033331-3102021232010220-0303101130112303-1231231011111111-1333333323232331-3201222212222220-1131201122001211"></a>

<a id="canonical-0220313012303100-2332031300111332-2120303010302013-3320120013202210-2310203102121100-1003000302122122-2323133032331112-1221020123111323"></a>

## digest property — sha256_digest / 221233030030 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-3230202200230030-2321320132013221-3300232330200333-3101330130130022-3101132100202313-1031202111131111-2321203103322110-3212230013101200"></a>

## Next pages — sha256_digest / 221233030030 / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2210121320301100-3213010231111311-0223132131103032-1030023212322202-1222200221131333-3100002232003002-3210211231131031-3123013311020031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232130323121212-0120122300221023-2202101022013020-0111003133210211-1202031022331102-2100132112110303-0011230030131131-1212333202333023"></a>

## primary.default_rr_set_group.ds_record.values.sha384_digest — sha384_digest / 332223332221 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303)
- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312)
- primary.default_rr_set_group.ds_record.values.sha384_digest

<a id="canonical-2030122233123310-0212122201203030-0330300330233113-1012010311212221-1032012133301211-2322013320000303-1330230301312032-2122102201210333"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231332132303320-0232101222310011-0031223020323110-0132000333230031-1130110000020133-3033131102210311-1120332320130332-3021212132300132"></a>

## Direct properties — sha384_digest / 332223332221 / 3

<a id="canonical-1101302331200220-2221110300012310-1122321000233133-3321121322212313-0110323200132312-2302203101220021-2323032100233233-0223111332330230"></a>

<a id="canonical-2331330331010123-3221203202133203-3212221001222213-3023103013211110-1101032211110102-1020320300033122-2323231011001030-2030112201233300"></a>

## digest property — sha384_digest / 332223332221 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-2001023322100100-2320203312332002-2012330120020120-2000112212332100-3120031102302133-2103210332020312-2223232210311211-2333110100131232"></a>

## Next pages — sha384_digest / 332223332221 / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3100001201120203-0312033322023013-2332030222311021-2122332202011202-1210221203131033-3222103132311001-0030203223010032-3022320213132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003231122220211-3133012223301101-0322323310023002-0130233130001123-0202021203303322-0212022223133022-3221203300031233-2012322021221331"></a>

## primary.default_rr_set_group.eui48_record — eui48_record / 223302123333 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.eui48_record

<a id="canonical-1231213302032020-1211213223312233-0233303332122300-3302330213013023-3010311130120300-1112221013032201-0001032011202133-0231203302203202"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui48 record.

Upstream description:

DNS EUI48 Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
```

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
eui48_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012031132212321-3012101302202220-3013211212030000-1103331012111232-3202030100300011-0001002222312122-2032230200000023-3312132003310302"></a>

## Direct properties — eui48_record / 223302123333 / 3

<a id="canonical-1012100330302012-0202201221210022-2221103202122300-3130110113030033-3021012203320133-0233232220033102-3310101030113001-0102131212231112"></a>

<a id="canonical-2012132031132220-2111112321133132-3312222300020023-2121101123000323-0110323001101023-0203312001001302-1130011100230032-1203211121123133"></a>

## name property — eui48_record / 223302123333 / 4

Type: `"string"`. Optional.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-3201112200020032-3123032103131101-1033011212121330-1312202003131323-0320123031331112-0020031212003120-3300330303011223-3111022322202313"></a>

<a id="canonical-1112100000123012-0002333233102133-1333123111111232-0310230202032212-3132321021101202-0122200020022133-0223233032133202-0100222111301323"></a>

## value property — eui48_record / 223302123333 / 5

Type: `"string"`. Optional.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Upstream description:

A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-0112113301120100-2210202032333323-3232130221133231-2013101202001030-0200003102002113-3113232002220030-0210122301223223-3033101233012203"></a>

## Next pages — eui48_record / 223302123333 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0310232010230002-2032333311200131-2210000030122303-1120200331023020-3101110012303221-2302222310112000-2232130122311323-0203010122113311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232013220223120-3033303010111331-2011321232202322-2101320330111333-3100121112320113-1120002100301332-2022033302031032-2301032111023122"></a>

## primary.default_rr_set_group.eui64_record — eui64_record / 030232312230 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.eui64_record

<a id="canonical-3223320120003011-3103320130031032-3311330233013221-1313223103010030-1113212203021121-0022012230030112-1113312132032223-1322223101102132"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
```

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
eui64_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202011330031020-1121110221212320-3231201232233320-3100122220213201-2333212121223132-0031133300120022-1200320332213233-0320303111022321"></a>

## Direct properties — eui64_record / 030232312230 / 3

<a id="canonical-1110123123300000-2102133000131311-2223001311203032-1012132033120023-1113322230311331-3332032332102122-1302103333220210-3311012123110000"></a>

<a id="canonical-3102010213321203-3322022022232030-3010300103123110-0012032013301101-1222233003121022-0102111231311110-2031323331200031-2020121311000331"></a>

## name property — eui64_record / 030232312230 / 4

Type: `"string"`. Optional.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-3030022232113003-1122010013003310-0022001100231103-3021222011200202-0130033322011032-2331223311112000-3231101330333020-1213331123313213"></a>

<a id="canonical-2002131210213102-0122322231211033-2322122122103310-0302102220012212-2111102212021033-3033002201011010-3300230313130102-3211210212303300"></a>

## value property — eui64_record / 030232312230 / 5

Type: `"string"`. Optional.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(23, 23),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-2120211113010213-3301222322122103-2203312030133320-3312003122223112-0200313332320113-2211100103121111-3003022323233020-2031131122100322"></a>

## Next pages — eui64_record / 030232312230 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2000023011031312-2003213123011132-0300331211333233-1003131020211202-0213013000200212-3132103320112330-2130111300321023-0221122033031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123232021311300-1313321221012121-3132010030013312-0202033231111112-1123212221313201-1100013010301021-0000102233001222-1310200212321201"></a>

## primary.default_rr_set_group.lb_record — lb_record / 000310232030 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.lb_record

<a id="canonical-0123002111321001-2130320220120210-1321310003303003-2100203031031100-2233221131213003-3032110020220011-1013001123223221-1101033233232310"></a>

Type: `"object"`. single nested block, Optional.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

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
lb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002001113321203-3123133223110102-0310220113100001-3322201223232012-1320022132101212-0010332020313213-0103021310331221-0321130000000102"></a>

## Direct properties — lb_record / 000310232030 / 3

<a id="canonical-1202012023213130-2103102332233010-0230120330112132-1212223011111002-2022131301321010-1003300222312330-0103010213100322-1201001121033001"></a>

<a id="canonical-0123220003221110-2030333232132023-1013213331010222-0011202030320112-0132122231013020-0233131111020201-3321112320021301-2120303320230103"></a>

## name property — lb_record / 000310232030 / 4

Type: `"string"`. Optional.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](resources--dns_zone--reference--group-002.md#canonical-1130010332130022-0221033112200030-1231013030221321-1331113132211300-1132223203031221-0333312030123323-2123213012013111-0031011201002313): complete subsection reference.

<a id="canonical-2300213312013330-0202002120130022-0220330021013110-2030110021033332-2220010210313220-2003120300313132-2313030113310130-2232113003013010"></a>

## Next pages — lb_record / 000310232030 / 5

- [primary.default_rr_set_group.lb_record.value](resources--dns_zone--reference--group-002.md#canonical-1130010332130022-0221033112200030-1231013030221321-1331113132211300-1132223203031221-0333312030123323-2123213012013111-0031011201002313)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1130010332130022-0221033112200030-1231013030221321-1331113132211300-1132223203031221-0333312030123323-2123213012013111-0031011201002313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131001300103000-3110132102133232-0033011323023221-1130133303220211-2100033211130102-0033212130303211-1023003313322332-1320230033021313"></a>

## primary.default_rr_set_group.lb_record.value — value / 232031230302 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-2000023011031312-2003213123011132-0300331211333233-1003131020211202-0213013000200212-3132103320112330-2130111300321023-0221122033031030)
- primary.default_rr_set_group.lb_record.value

<a id="canonical-3323100013132323-2203130020313231-2231013033222222-0303012302232332-1033302010331212-3320121211101001-0313002232032000-2111101103302032"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332330130211020-3020111212332330-3102101120210221-0030213020003032-1101211000213232-2020021032321102-3311022223220212-0102223232332212"></a>

## Direct properties — value / 232031230302 / 3

<a id="canonical-0210222223002130-2230001023300130-0101330120013331-2003001200212201-2330301013202032-1020331323121313-2330111010203310-2330210310031302"></a>

<a id="canonical-1303031112223022-3102030202032131-0202332000013032-3012013032302322-1100112322002123-1230322013233211-3210223012212322-0021323110332300"></a>

## name property — value / 232031230302 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2131202322332212-3002101101303303-0113312022321123-0022323100111111-3222131232310203-3312003313113011-2211113303301103-1310021021213013"></a>

<a id="canonical-0201020013330030-1103210301110300-3121223331001212-0002011133121333-0021312000001103-1312031111320312-2230330132003123-3230022132132200"></a>

## namespace property — value / 232031230302 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2311232221230032-1222033332001211-3213122203121101-1111312321031020-1230020213322220-0203213030132232-2020220311001121-3131322201223123"></a>

<a id="canonical-0310023233310121-0231102301223032-3031310001122232-0323112131111312-0023100123101320-0020211000010212-0322302213002013-3200100112202113"></a>

## tenant property — value / 232031230302 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0133021011120021-3202211323223113-2021210132323330-1231023031111220-1030321033001122-0322303101310102-2321330303101312-3111321330313022"></a>

## Next pages — value / 232031230302 / 7

- [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-2000023011031312-2003213123011132-0300331211333233-1003131020211202-0213013000200212-3132103320112330-2130111300321023-0221122033031030)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0001331322032313-1331233300221231-2021032123331122-1313301111122120-1113123132222021-0013311333002210-3021113111120010-3121010323223333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320001301011213-3120311233302313-1201212230131332-2113103122323202-0022331022310310-2120002332112030-0131232123302001-1213313033003031"></a>

## primary.default_rr_set_group.loc_record — loc_record / 121203332002 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.loc_record

<a id="canonical-2132111113030021-3321322230122210-2330022320201032-0022111313131111-3112122321323003-3321113201120112-2331221333120033-0232020030113010"></a>

Type: `"object"`. single nested block, Optional.

DNS LOC Record. DNS LOC Record.

Upstream description:

DNS LOC Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
loc_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310132223232010-2010011123002202-1201313002230100-2111311300321100-2331010123301201-0113221031132230-2213230123223032-0123321023321100"></a>

## Direct properties — loc_record / 121203332002 / 3

<a id="canonical-0012123332333200-2002113200301200-3333120212310101-1112202231311122-3313130131110103-3032330100233230-3031333110113100-2122330012023202"></a>

<a id="canonical-3003102231021003-1201332111232013-3122010030033330-3323211213323313-0221100301012322-1120102020033030-3221023212302211-0221100100032312"></a>

## name property — loc_record / 121203332002 / 4

Type: `"string"`. Optional.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-2112121111131001-3303232101130313-1333110202023323-2101112231121012-1103313011133322-1033130001232233-0232023330011221-1223033200111230): complete subsection reference.

<a id="canonical-1310113030022012-0113321320301110-2333222123303012-1302201231021233-2212132121122100-3111322211321132-3330112313130223-2133212322030111"></a>

## Next pages — loc_record / 121203332002 / 5

- [primary.default_rr_set_group.loc_record.values](resources--dns_zone--reference--group-002.md#canonical-2112121111131001-3303232101130313-1333110202023323-2101112231121012-1103313011133322-1033130001232233-0232023330011221-1223033200111230)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2112121111131001-3303232101130313-1333110202023323-2101112231121012-1103313011133322-1033130001232233-0232023330011221-1223033200111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031221320001033-0022320101121320-1120021020202102-3130210022123102-2031200123202123-0330031220233002-3200310113013031-3223331323233321"></a>

## primary.default_rr_set_group.loc_record.values — values / 323022003113 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-0001331322032313-1331233300221231-2021032123331122-1313301111122120-1113123132222021-0013311333002210-3021113111120010-3121010323223333)
- primary.default_rr_set_group.loc_record.values

<a id="canonical-0000303313130103-3211013021131123-2212112102013332-2313032301323331-2113023333303021-2122000331232010-1033313323030300-0231003010101302"></a>

Type: `"object"`. list nested block, Optional.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("altitude",
    "latitude_degree",
    "longitude_degree")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230300022122321-3102301121101213-0311030321110322-3310121100331213-3232302131220322-3332222312200102-2312022021100212-0321200121131000"></a>

## Direct properties — values / 323022003113 / 3

<a id="canonical-2131202112103221-2211310101233311-3013113321130321-3223213123102001-1223001232012103-2111100230303110-2011211320302332-0030000332201231"></a>

<a id="canonical-1331110213012320-3132311211200202-0203220021123230-1313221302013200-3010330102001221-3201203132033130-0011133211111002-0112303031233223"></a>

## altitude property — values / 323022003113 / 4

Type: `"number"`. Optional.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0223310311003231-2130112230301222-1200002110121021-3023210233223033-2013112222210132-3110100112113332-0022003203112030-0013132220111321"></a>

<a id="canonical-3310102010131301-3210100123210211-2300323020123233-0020333230122213-1102233332212013-1213030333300131-3211303032213311-0310102013323121"></a>

## horizontal_precision property — values / 323022003113 / 5

Type: `"number"`. Optional.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-3101333230031331-3203232301103020-0123322112112103-3013313202211231-2002101312323213-2033001121131010-3313131200130212-1330321222233011"></a>

<a id="canonical-3031212112222331-1000003122221211-0233013213333333-2100331132222303-3030003013232132-1010023013103102-3331230330321321-2310331102133220"></a>

## latitude_degree property — values / 323022003113 / 6

Type: `"number"`. Optional.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 90),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3010322301212100-3221131321000232-2003311301111310-0231101220002001-1030003102211020-0301203102230320-3221023213233001-1002210230231013"></a>

<a id="canonical-0302013330223032-3001301313020103-0300212220130320-3331231313230232-2102320333220031-0021200011210132-1310011011331013-0200102321100032"></a>

## latitude_hemisphere property — values / 323022003113 / 7

Type: `"string"`. Optional.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("N",
    "S"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3003200010121011-3122020100020232-3120300303303100-1300330210031023-1101122321033011-3320131033011310-1113300233220011-1130013003311130"></a>

<a id="canonical-2302213112033022-1103311032133303-3100120221221222-1231123100122223-0313321012031333-3321133202113230-3013111123002132-1031100100320110"></a>

## latitude_minute property — values / 323022003113 / 8

Type: `"number"`. Optional.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-1132311121131210-0032113300321100-2230002011302100-1312130031130022-0023131012023001-3011231130111301-1310210111302010-1112100100223123"></a>

<a id="canonical-1231301212101220-0122213020231202-0220202230221302-2311221032111301-2231321132033113-0203300130113021-2131010033312210-3122213213212233"></a>

## latitude_second property — values / 323022003113 / 9

Type: `"number"`. Optional.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-2001302000003023-0110113301110220-3022310200322013-1322311000011013-2203323331133222-1332122010032033-1022302221123222-2310221233123231"></a>

<a id="canonical-0020003313201313-3302002211211230-0102302113022213-0201200322001312-1301201032032323-3130231121103223-2112000222101331-1212301211202102"></a>

## location_diameter property — values / 323022003113 / 10

Type: `"number"`. Optional.

Diameter of a sphere enclosing the described entity, in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-0123032002122113-1101033300112321-3102333111012213-3111120331331330-0211113310100321-1201301111002221-2113201211123032-2231310312113203"></a>

<a id="canonical-0300120212301023-1231000310101032-1010022102221132-2121132032013303-2032100033212133-3321323033330032-0123112100203312-2303221311321212"></a>

## longitude_degree property — values / 323022003113 / 11

Type: `"number"`. Optional.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1033200313203132-3302301021222000-2012320320112210-3201220231010301-0013323120103203-1232023331020331-2301011211131303-2013000022210032"></a>

<a id="canonical-3022032313313213-2301203201001201-2032020131122011-1222202311323321-3221000033333313-1212031113112101-3212102002003330-3230100200121201"></a>

## longitude_hemisphere property — values / 323022003113 / 12

Type: `"string"`. Optional.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("E",
    "W"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3101310132203301-1232323301110103-1023221213131013-0330301121231202-1033222223012023-0221300103211313-3200213130320330-0201323023322231"></a>

<a id="canonical-1103102032011100-2203022122200120-2313030333000000-2303220033200113-1012303132223100-2200021020102012-3101302201203222-3000100233311121"></a>

## longitude_minute property — values / 323022003113 / 13

Type: `"number"`. Optional.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-2023230311100231-3322110003113122-1033331333231311-0332023131230210-2233133003101311-2033101222223113-3110013203002202-0322120201123130"></a>

<a id="canonical-2011132131313210-2013310301210331-2211003233022303-2212333001211033-2033222111103031-2223101000302031-3012300122222000-0300132012210102"></a>

## longitude_second property — values / 323022003113 / 14

Type: `"number"`. Optional.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-2303210211122011-3132033003000301-0102311002020001-0200122221213003-1021221132220102-1322301213122313-2220221303100310-1220013111321300"></a>

<a id="canonical-0130231231211121-2310203333103100-3312213222122033-3100332002110230-2230213103221212-0302133313222033-1013012302201111-2103000212012023"></a>

## vertical_precision property — values / 323022003113 / 15

Type: `"number"`. Optional.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-0100323222320021-2211322023331310-0022233321022303-0112030012232202-1123001033201122-3111312100221033-3010211201223032-1001133210031121"></a>

## Next pages — values / 323022003113 / 16

- [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-0001331322032313-1331233300221231-2021032123331122-1313301111122120-1113123132222021-0013311333002210-3021113111120010-3121010323223333)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230032010113001-1102301231000013-1300013210320113-2012002111202210-3202022210230012-2100321213231313-2103300321333231-2130222331033330"></a>

## primary.default_rr_set_group.mx_record — mx_record / 211332222003 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.mx_record

<a id="canonical-3031022012121010-1033000003212220-1333213133232001-2301203330223011-1202332101303120-0230200230301011-1311002130031131-0113212112300021"></a>

Type: `"object"`. single nested block, Optional.

DNSMXResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
mx_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231033012101333-3322101210032031-3131213310120111-1102230331301133-3321231030312321-0301000203101331-1130220212102332-2010101112132221"></a>

## Direct properties — mx_record / 211332222003 / 3

<a id="canonical-1321002202112310-2312002321302110-1130310200333323-3211330321331112-1031301011303103-1103133023220232-0302202132003202-0230330012010013"></a>

<a id="canonical-2003101312312312-2312231113131231-0100102332001303-1303223300002131-2213120303011033-3122331123121322-3133310102233103-3300333320311322"></a>

## name property — mx_record / 211332222003 / 4

Type: `"string"`. Optional.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-2132212202122320-2030321100131003-0133003301133102-0220221112310121-3122231332022001-3032232202002320-0113111230100302-3313121002202311): complete subsection reference.

<a id="canonical-1320223311213203-0120102233332010-0130032313001212-3032122301233113-1233023020200311-3331221133313223-0233303011210133-1301103220232012"></a>

## Next pages — mx_record / 211332222003 / 5

- [primary.default_rr_set_group.mx_record.values](resources--dns_zone--reference--group-002.md#canonical-2132212202122320-2030321100131003-0133003301133102-0220221112310121-3122231332022001-3032232202002320-0113111230100302-3313121002202311)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2132212202122320-2030321100131003-0133003301133102-0220221112310121-3122231332022001-3032232202002320-0113111230100302-3313121002202311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311000231202133-1320022022133002-2101231302110033-2000030201313130-3132303331202021-1000333333301333-1110012023020231-0013323102030301"></a>

## primary.default_rr_set_group.mx_record.values — values / 103230121002 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312)
- primary.default_rr_set_group.mx_record.values

<a id="canonical-0003030310133033-1301312201120100-1102301001211011-1101203121013200-3232121203223313-3200201303223333-3212301101202311-2110102332330230"></a>

Type: `"object"`. list nested block, Optional.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221112310133230-2330000110223312-2010222001102211-1003330123221003-2211222023100020-2203130103323023-2231331300100021-2332111203102323"></a>

## Direct properties — values / 103230121002 / 3

<a id="canonical-3212223332222301-2111210213031321-1302011002231203-2203220223311120-2233203000210201-2221000222013101-0013120102130102-2020000013302301"></a>

<a id="canonical-2123300001311000-0321011301322122-0213323311211301-1303222203010022-0111232023030310-2221203313002223-1111312130012333-3230310321212132"></a>

## domain property — values / 103230121002 / 4

Type: `"string"`. Optional.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0301021110310111-1322301311102201-1231000121230232-1333010102002301-2202120233313002-3303212202312212-3030323011221232-2032301101201300"></a>

<a id="canonical-1333122330011122-3023101313322032-3322201131000130-3123021101120103-3112023303311323-0033100133312113-2110110222010020-2103211310332201"></a>

## priority property — values / 103230121002 / 5

Type: `"number"`. Optional.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3013110000202231-3010113310030132-1333201223033300-3003320021003223-2012322120331303-3230312103231113-3101133221300312-0133131103100330"></a>

## Next pages — values / 103230121002 / 6

- [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0221201201321311-3033210312022203-0010211233110120-0301032022010211-2322320001313111-2002020120203123-1333000323301120-3330303320333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301313031010320-3100310100312111-3311113222131313-3030231202100013-2123302102211113-2033131303332012-3302201111322200-0031123233203133"></a>

## primary.default_rr_set_group.naptr_record — naptr_record / 100201223133 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.naptr_record

<a id="canonical-3302201020210130-3022011223202021-0033303101220201-0002010112333122-0031033012232221-3313032200103232-3102321021201211-1101130203111310"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
naptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100321210230311-2020303222011301-1133110333010033-0113100123210120-1021222301311133-0200211000313233-0132103220210010-0221201321103013"></a>

## Direct properties — naptr_record / 100201223133 / 3

<a id="canonical-3303200222203311-3322303122022321-0310111001311232-1311233100222002-2000210202131011-3322313011110023-3133012122231120-2120030332332222"></a>

<a id="canonical-0120311113313320-0122132222221103-3033311112322210-0030311232333031-0312120020321101-0332230023030301-1022110132002222-3021201131312001"></a>

## name property — naptr_record / 100201223133 / 4

Type: `"string"`. Optional.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1111030123330233-3012003313010210-1310201033210001-1202023220222030-2122100010113323-2232130003103221-0031211122130333-0020302323221311): complete subsection reference.

<a id="canonical-0123222121030200-2230223122013221-3303301033033131-1232131020301110-2012122211121333-3010030020023333-0200120010201312-0220003202020232"></a>

## Next pages — naptr_record / 100201223133 / 5

- [primary.default_rr_set_group.naptr_record.values](resources--dns_zone--reference--group-002.md#canonical-1111030123330233-3012003313010210-1310201033210001-1202023220222030-2122100010113323-2232130003103221-0031211122130333-0020302323221311)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1111030123330233-3012003313010210-1310201033210001-1202023220222030-2122100010113323-2232130003103221-0031211122130333-0020302323221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010233333033332-0310230310301002-1303002232123202-1323000330111201-1232000023012013-3211200222201112-1101031130321020-2133333210012312"></a>

## primary.default_rr_set_group.naptr_record.values — values / 330201303312 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-0221201201321311-3033210312022203-0010211233110120-0301032022010211-2322320001313111-2002020120203123-1333000323301120-3330303320333302)
- primary.default_rr_set_group.naptr_record.values

<a id="canonical-0012133211132101-0222112211233003-0202312222331133-1111020032121221-1311303333202031-3112101213011201-2030200321213312-3020100000132311"></a>

Type: `"object"`. list nested block, Optional.

NAPTR Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("flags",
    "order",
    "preference")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202132021311320-1111321010323002-0021103113302322-0131320330002211-2202222122313221-2330132002332013-3113221202003000-1123003012010022"></a>

## Direct properties — values / 330201303312 / 3

<a id="canonical-1033021233210122-3200232001131212-0130330200111101-2002003333123000-3310031000110031-3230321033232230-0130211012022123-3303002312103123"></a>

<a id="canonical-1203302011203323-3102030021212031-0303201030202331-3102300130222203-0132130220023103-1311300312010132-0132013023303321-2312220030230200"></a>

## flags property — values / 330201303312 / 4

Type: `"string"`. Optional.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-0003322302231222-1002002201313010-3201230211230220-1223201122320320-3311330200200203-0000121233330012-0311232032101331-3123313013332232"></a>

<a id="canonical-3322230013113012-3033120023022231-2002220300133201-1313210201213320-1132311322201012-0203100102031131-2331132011203222-1332102010300320"></a>

## order property — values / 330201303312 / 5

Type: `"number"`. Optional.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0203233120120121-1223303131330102-1223232310132002-3301221113103100-3312233033302311-0011020332332101-0011323233301132-3221221100312303"></a>

<a id="canonical-3202312033223000-1122023220102013-0301201232001320-1023330000031202-1222220133013103-3033221112233303-2201222013120221-0123120120031020"></a>

## preference property — values / 330201303312 / 6

Type: `"number"`. Optional.

Preference when records have the same order. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1133101022121030-3222111130232323-3003312223130130-3000111311111320-0102021122022012-3123102310331021-2331310210210022-1100132112013120"></a>

<a id="canonical-2200302133200300-2212002213023203-3130033132101333-3210210211101233-2221102001220120-3230031022203320-1131212310102003-3003012200333011"></a>

## regular expression property — values / 330201303312 / 7

Type: `"string"`. Optional.

Regular expression to construct the next domain name to lookup.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1303002220120012-3033311223202121-2210031032110320-0220031103301301-1110120230031311-0310212111211132-0101022331313302-1221303110022131"></a>

<a id="canonical-0111301110221323-1201012300231033-1030020112232002-0311312321110200-0130111300002000-2221203113103201-3100211312223033-1302120002131123"></a>

## replacement property — values / 330201303312 / 8

Type: `"string"`. Optional.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0100322000333330-1013302333103132-2202202223102312-1101112312031213-3310021222110212-1220033222200200-3132030223101200-0120130112301221"></a>

<a id="canonical-2202331122213232-2100103102221123-1331233111201333-1311320110120221-2103030231201100-1231332300113102-1000331311012203-0033201330201321"></a>

## service property — values / 330201303312 / 9

Type: `"string"`. Optional.

Specifies the service(s) available down this rewrite path.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-0230320033222310-1002331022230331-3102003102222012-2310331203313000-3311112031212213-3233231212212033-0231103113113333-1030200001232330"></a>

## Next pages — values / 330201303312 / 10

- [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-0221201201321311-3033210312022203-0010211233110120-0301032022010211-2322320001313111-2002020120203123-1333000323301120-3330303320333302)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1320012000010221-1013101202011330-0222031102033131-0333201222231300-2310103201032220-3222020113330211-0230001303312231-1302222302130302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232133033111132-1031011000322101-2233021312212003-0110021101033033-0322233313312030-0310231011123013-3003213001213121-1011132001203321"></a>

## primary.default_rr_set_group.ns_record — ns_record / 003210211300 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.ns_record

<a id="canonical-3133323100030001-3323320220032012-1023332123310223-2133233223111122-2110331310221011-3223131111103203-2022210101000332-0012033013311031"></a>

Type: `"object"`. single nested block, Optional.

DNSNSResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
ns_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111023131013111-3031223021023031-1010120233000133-3320202231012111-1020001303030020-0002332311101032-2302223222312020-0201222223131312"></a>

## Direct properties — ns_record / 003210211300 / 3

<a id="canonical-2033003111331232-2100321300133202-0112031132303032-0321230200103202-3333010122311223-3320231030313320-2300021021310111-2111212112023110"></a>

<a id="canonical-2323213223202111-2030130233232022-3102211023112132-0013122012230030-1301310003101312-3133130211133303-0300301100022021-0020202122333320"></a>

## name property — ns_record / 003210211300 / 4

Type: `"string"`. Optional.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3113333313223202-3322202032213010-1202233000332200-3222103322113102-2222133303301233-0320203210302131-3231111130301002-2312123022210102"></a>

<a id="canonical-1111320313113303-3231302232310020-3231332100023002-0003333002113332-1032331201332003-0032222120312010-1120133013330221-3301110311110101"></a>

## values property — ns_record / 003210211300 / 5

Type: `["list", "string"]`. Optional.

Name Servers. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1013102303110031-2122330101302223-0013220120011113-3113013210300332-2103112220023211-3222020102301101-0111033231201323-0012023232300011"></a>

## Next pages — ns_record / 003210211300 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3030100213211033-2113011011200322-2233222032221031-1220020031011201-1132132112023130-3211231011202332-2123323330002213-2200232102012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332311012223110-2123121003300022-2223001210303000-2331010212303231-0231001223130331-3231021110031101-0121100010232013-3330133020310331"></a>

## primary.default_rr_set_group.ptr_record — ptr_record / 003332110011 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.ptr_record

<a id="canonical-2000130221010012-1133232130332203-2023123012103123-1213333232011033-3021323331322132-0220320211113333-2311120321002201-0202020210201231"></a>

Type: `"object"`. single nested block, Optional.

DNSPTRResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
ptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030300311002122-0032302212233012-0233331213313213-1332303032231303-0310030302101022-0212003210321031-1102322103122003-0202310111302031"></a>

## Direct properties — ptr_record / 003332110011 / 3

<a id="canonical-1102100003332030-3003322001321331-2232010102001313-3223110102112111-1320323213032321-1132212230311002-3130231300002110-0132033303303000"></a>

<a id="canonical-2032030312301203-1310113331003110-2003133111211032-3023000123022000-3311020113212320-0131213230301213-3330123331212210-3100231103123012"></a>

## name property — ptr_record / 003332110011 / 4

Type: `"string"`. Optional.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0212323031323213-3310023311233130-3013102103220130-2030220310203210-3221031003330200-3230100102132031-0110120031033122-3022123022013230"></a>

<a id="canonical-3012033032222330-0311331321123031-2001021210211112-0132013023030021-1130321021321021-1021200300231031-0201230012111100-1212201202103211"></a>

## values property — ptr_record / 003332110011 / 5

Type: `["list", "string"]`. Optional.

Domain Name. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2202210303313232-3203332221021211-1003100030321100-2012321022003301-1013033202220001-0110100211230003-2211110133200221-1211311022033312"></a>

## Next pages — ptr_record / 003332110011 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0113001230120133-2303121012102031-1200311122002232-3032331232033323-1200010123122221-1212320123121123-0013213201323320-3313313211120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033203220331123-2013010000311112-1102323331331032-2121203130200131-2010002203123311-1113001113220021-3010303232223223-3030210302012210"></a>

## primary.default_rr_set_group.srv_record — srv_record / 321302230113 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.srv_record

<a id="canonical-3232132023001031-0110132101031030-2312320323010000-0302113021120313-2200303232110120-0032332303330320-2200231003220012-3002111130330010"></a>

Type: `"object"`. single nested block, Optional.

DNSSRVResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "values")}
```

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
srv_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201010222213311-3231223220300203-1332230031123112-2303212332020312-1203122030223321-3010103111220320-2131220122132002-1003230222132113"></a>

## Direct properties — srv_record / 321302230113 / 3

<a id="canonical-2021022302320110-0232311121203113-0131321030222332-3321300123301022-2202212013111310-3002012230220200-3010121102310002-1211021110021100"></a>

<a id="canonical-1022302121222023-1211200003321303-1210212302232031-1133133013212322-0010122311332120-1001122023123311-0313013231220001-2013121221312223"></a>

## name property — srv_record / 321302230113 / 4

Type: `"string"`. Optional.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-0220200230321313-1301130010021321-0111211323133021-2233302021322332-0311101110010303-1222201032221100-1000122112231033-2103231133330232): complete subsection reference.

<a id="canonical-3322300300001202-3302330130121113-1311011212133111-0030330232313212-3330310031301023-2221122301131301-0133110323202233-3223121300330003"></a>

## Next pages — srv_record / 321302230113 / 5

- [primary.default_rr_set_group.srv_record.values](resources--dns_zone--reference--group-002.md#canonical-0220200230321313-1301130010021321-0111211323133021-2233302021322332-0311101110010303-1222201032221100-1000122112231033-2103231133330232)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0220200230321313-1301130010021321-0111211323133021-2233302021322332-0311101110010303-1222201032221100-1000122112231033-2103231133330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321121013031301-3333132000111030-2230321310013030-1221322211321003-2313320120203002-3030131030111033-0000231320230321-1301233113332223"></a>

## primary.default_rr_set_group.srv_record.values — values / 002012233303 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-0113001230120133-2303121012102031-1200311122002232-3032331232033323-1200010123122221-1212320123121123-0013213201323320-3313313211120233)
- primary.default_rr_set_group.srv_record.values

<a id="canonical-1233110110121133-0211332212112022-0220333223110020-0222101213000123-1223321332033030-1212201202333331-0222210032003303-0133220311010220"></a>

Type: `"object"`. list nested block, Optional.

SRV Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101331133301333-1030111000213221-0302310033011120-0020121213021113-3131032200013011-1020030231100301-0013122212201002-1102021131101311"></a>

## Direct properties — values / 002012233303 / 3

<a id="canonical-3020210212331312-2113033213301003-3332030210023002-1222331100023311-3213033310033211-3020101212033020-3002301000121233-0023302121210313"></a>

<a id="canonical-1330303021221322-3332031022230010-0132130233221123-2300133022023330-1210132330020013-2232012132031302-2323321202131310-3220233323030013"></a>

## port property — values / 002012233303 / 4

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0212122213022002-1102130330230022-2113103310311202-3022213003321111-3031122031223210-2302032331122032-0033223320030102-3021033322001002"></a>

<a id="canonical-2212002001232202-1210001330202100-3210322121023021-2313120222113212-0323030103020231-3012113111310020-0323212331312023-3130022202320332"></a>

## priority property — values / 002012233303 / 5

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3223322100303321-0031100122012000-0231120301220323-3101323020211220-3321023031332030-2000233003301022-1012301232212321-0223131202102230"></a>

<a id="canonical-2101031221222201-0311100101110323-3221300210113202-2233231032113323-1320023123300013-2220131130011013-0003200012032213-3213302001311223"></a>

## target property — values / 002012233303 / 6

Type: `"string"`. Optional.

Hostname of the machine providing the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-0223320020032303-2313320021010002-1100112012113203-2222112110223313-0220113331223131-1330112001302113-1102223203003133-2303132101301102"></a>

<a id="canonical-2000020231013220-0310032320013023-0011010313010120-3203220103131321-1213310130223010-1221032011121003-0133230310231331-0132221023222302"></a>

## weight property — values / 002012233303 / 7

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2300022210222211-2110132130023100-1323213203323330-3303120332230213-2212010013201101-1313100332321323-1101311311111022-1022233103221233"></a>

## Next pages — values / 002012233303 / 8

- [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-0113001230120133-2303121012102031-1200311122002232-3032331232033323-1200010123122221-1212320123121123-0013213201323320-3313313211120233)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012022302233122-1112222301103030-0133133123002302-0133323030300321-3213200311001023-2100223021010122-1100311332223311-2300330030323130"></a>

## primary.default_rr_set_group.sshfp_record — sshfp_record / 202311332121 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.sshfp_record

<a id="canonical-0033310300233313-3001201311021021-3123210133213011-0313232002021102-0111202110303313-0230133031020323-2110230212303300-3320313013021301"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sshfp record.

Upstream description:

DNS SSHFP Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
sshfp_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332310313302002-2200201112102030-1000133020122213-2230100313000232-3302312311221310-3300232011310120-2301210233132301-3322102020000313"></a>

## Direct properties — sshfp_record / 202311332121 / 3

<a id="canonical-1312323000120211-3122313122033001-2301100232310012-0010001313331121-1103323211133023-1111210113022110-0113201100200021-1323210100121211"></a>

<a id="canonical-2002333210122121-3230303121202020-3003010011023221-3211333312122313-3331222211232302-3201203302313232-3010133223312310-3112331210233321"></a>

## name property — sshfp_record / 202311332121 / 4

Type: `"string"`. Optional.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212): complete subsection reference.

<a id="canonical-2033203000320012-1233232112011223-3311311210011110-1200220203320132-2233310331100101-1031030000033103-3313323110003112-2330120331221120"></a>

## Next pages — sshfp_record / 202311332121 / 5

- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222222210111320-0012013103231030-1112001111121300-2000023322221032-0102012210132233-2200013002032013-3202233331113200-3222232331203213"></a>

## primary.default_rr_set_group.sshfp_record.values — values / 333110111013 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- primary.default_rr_set_group.sshfp_record.values

<a id="canonical-0323200310303211-3133202223300230-3333231133230112-0322212032230011-1120330010031221-3103010302313102-3310223021330202-0222321323003333"></a>

Type: `"object"`. list nested block, Optional.

SSHFP Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sha1_fingerprint",
    "sha256_fingerprint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231032120002133-0022202313322200-3231123200302100-0331311230101110-2222030103133120-2201020202121300-2102000222012112-2333232211213222"></a>

## Direct properties — values / 333110111013 / 3

<a id="canonical-0100312300001122-1100322030101033-3013022100332131-0233211220102221-3132132323131321-3013000222230301-1021002323200020-1320101113003203"></a>

<a id="canonical-3023133311012032-3201002230111132-2122122031312302-3001332000302001-2110010231331022-1210221301021123-0302202030101313-1001010130311311"></a>

## algorithm property — values / 333110111013 / 4

Type: `"string"`. Optional.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Upstream description:

SSHFP algorithm value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM

&#8203;- RSA: RSA

&#8203;- DSA: DSA

&#8203;- ECDSA: ECDSA

&#8203;- Ed25519: Ed25519

&#8203;- Ed448: Ed448.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](resources--dns_zone--reference--group-002.md#canonical-2023231033131113-2110202123131020-1203111201302332-1313212001113301-3333113320331021-2020121001100123-0220312010123310-2000300312120110): complete subsection reference.

- [sha256_fingerprint](resources--dns_zone--reference--group-002.md#canonical-3103022332110213-3100031323212012-0221003020003232-0230303211030301-0033003030113002-0032113003232322-2323001210333103-1021113222213111): complete subsection reference.

<a id="canonical-1132221111112112-1132231321203032-3333230210020322-0100032223100321-2332100121223321-3000303302133021-2112311203311300-0333223003220302"></a>

## Next pages — values / 333110111013 / 5

- [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-002.md#canonical-2023231033131113-2110202123131020-1203111201302332-1313212001113301-3333113320331021-2020121001100123-0220312010123310-2000300312120110)
- [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-002.md#canonical-3103022332110213-3100031323212012-0221003020003232-0230303211030301-0033003030113002-0032113003232322-2323001210333103-1021113222213111)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2023231033131113-2110202123131020-1203111201302332-1313212001113301-3333113320331021-2020121001100123-0220312010123310-2000300312120110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321101202102302-1303110031311001-3110010132031131-2222032110202322-3300332221320300-1332313000123122-3223321013033300-1321201321321111"></a>

## primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint — sha1_fingerprint / 203130022110 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212)
- primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint

<a id="canonical-3300230102310032-0013311210030213-3111312011020222-2122123002032112-2221302013330301-3202101323320302-3302012213021222-2320211131333010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 fingerprint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
```

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
sha1_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011000303030030-0321313332123321-0011121130110102-2300130231203030-2111320220030332-3323301321200100-2100031223320200-0110132231230222"></a>

## Direct properties — sha1_fingerprint / 203130022110 / 3

<a id="canonical-3230330302020020-1213132222131323-0002010120323031-3233011323002013-2201200301032200-0121322201011222-1313221131123210-2001212211321122"></a>

<a id="canonical-3212230012330231-3331132320322310-1113222121000222-1002210202113300-2200222020321132-0031022012033123-0033331103301212-3000103312333222"></a>

## fingerprint property — sha1_fingerprint / 203130022110 / 4

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-3303211310210110-1220222322313033-0203331133302321-2133021222222123-2102110312120202-3020313113012101-2101320030233302-3013130333222113"></a>

## Next pages — sha1_fingerprint / 203130022110 / 5

- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3103022332110213-3100031323212012-0221003020003232-0230303211030301-0033003030113002-0032113003232322-2323001210333103-1021113222213111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303030123311212-3210203101120013-2312203332122322-0333031223103011-0000222121002200-3321032300031003-3000320213133321-0013321002322213"></a>

## primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint — sha256_fingerprint / 111221110112 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212)
- primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint

<a id="canonical-0123302233213302-3311232213203031-2013230021213311-3000122000201211-0020103220202111-1121230212202312-2201032200130123-3303320103313102"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 fingerprint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
```

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
sha256_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303213130323002-3123202220333032-1332331300303033-0003202300000312-1032220213231120-2103223213031201-3022300220122201-0121010110232111"></a>

## Direct properties — sha256_fingerprint / 111221110112 / 3

<a id="canonical-1302201230021103-1232233201132323-1030012333213200-1133320003102230-3323030011231103-1230022311102032-2000000122202223-3302112310002200"></a>

<a id="canonical-0210321201132131-2032213212331302-2001332033200013-2010321312211312-0213203332102322-0031002220130232-2222113011321103-1332232332203011"></a>

## fingerprint property — sha256_fingerprint / 111221110112 / 4

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-3322023101310323-2103012000222102-1230221033002233-1000222222311221-2310002231331213-1020202210033230-2121122101320201-1232332332213220"></a>

## Next pages — sha256_fingerprint / 111221110112 / 5

- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1110210111012102-2031202010320123-1132320121232332-2002023031230032-2022213032232233-1212213101223203-1103321022220000-0330112323013122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333113123231031-0323200223031110-3021133003001212-3211110031103220-1222301311031110-1123232202013202-2000211233203223-3020302111022103"></a>

## primary.default_rr_set_group.tlsa_record — tlsa_record / 122213023321 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.tlsa_record

<a id="canonical-1320110233003023-3320023020133213-1311023002130001-3013232233301000-2313223031331311-1202330011123222-3011223300111332-1022132230100103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tlsa record.

Upstream description:

DNS TLSA Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
tlsa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323313033011313-1320313010123001-2223010303323223-3113313111301320-1233312223000201-3310320330003123-0023233102231322-2313202333002030"></a>

## Direct properties — tlsa_record / 122213023321 / 3

<a id="canonical-1223313213332221-2321231333010013-2102333210131220-2031033201102231-0113101310133310-1212302101010313-2333033303031031-1302331113230213"></a>

<a id="canonical-0031222110303033-0011302210231310-3221203023201200-0012000232330133-3321012001233131-3021001130000101-1300020031003333-0330210033331330"></a>

## name property — tlsa_record / 122213023321 / 4

Type: `"string"`. Optional.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1103233213133131-2320013020321000-0101132130130100-1302121013001001-0101011131022333-1213332023301201-3301000123113300-0202231132001031): complete subsection reference.

<a id="canonical-1202331203211320-0032121312302203-1003031020312233-0332322333313300-2332221312333112-1300003030302203-3203302111032330-3331213220210203"></a>

## Next pages — tlsa_record / 122213023321 / 5

- [primary.default_rr_set_group.tlsa_record.values](resources--dns_zone--reference--group-002.md#canonical-1103233213133131-2320013020321000-0101132130130100-1302121013001001-0101011131022333-1213332023301201-3301000123113300-0202231132001031)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1103233213133131-2320013020321000-0101132130130100-1302121013001001-0101011131022333-1213332023301201-3301000123113300-0202231132001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033101133030313-2221020232202330-0110201211113023-3022002112002310-0200331200210001-1230133320303110-1112211223330113-1001232112013232"></a>

## primary.default_rr_set_group.tlsa_record.values — values / 123231100331 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-1110210111012102-2031202010320123-1132320121232332-2002023031230032-2022213032232233-1212213101223203-1103321022220000-0330112323013122)
- primary.default_rr_set_group.tlsa_record.values

<a id="canonical-2033223312322302-1020120221212021-1021023130230232-3102322210033211-1133210131303000-2202231002002201-2100002112220313-2113311310333233"></a>

Type: `"object"`. list nested block, Optional.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_association_data")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303023023213033-3012200032210013-2330303232001323-0323112031232122-1022220103122321-0312112323330201-0022212133122013-0133323332020103"></a>

## Direct properties — values / 123231100331 / 3

<a id="canonical-2202311121123332-2023231031233301-3303012300022301-0020033213311212-2210302332102100-2311212113013020-3223323000212021-1313100321202233"></a>

<a id="canonical-0210000321212133-3013020221121203-0101103012122201-3211101302113323-1121100012321333-0020233220020232-3323013310310212-0333300003321221"></a>

## certificate_association_data property — values / 123231100331 / 4

Type: `"string"`. Optional.

The actual data to be matched given the settings of the other fields.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0113022033122333-3313311131112032-3030022320121300-2333302321230022-1112323100213131-1232222311222310-2222211013312320-0032120012312110"></a>

<a id="canonical-1100323110332011-0213021023302020-1221100313213212-0222100123003323-2012011303100010-2331213103330120-3221101231030111-0221011213132013"></a>

## certificate_usage property — values / 123231100331 / 5

Type: `"string"`. Optional.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023202102001012-3112230200010101-2321201132100033-0011113231330002-1323022102101133-1232311132322203-2003302222023300-3112032022303023"></a>

<a id="canonical-0222131332310300-3231212110230321-3110102202010022-0121122030122233-1231103311033223-3323323321003032-0233012100102312-3313303301011221"></a>

## matching_type property — values / 123231100331 / 6

Type: `"string"`. Optional.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NoHash",
    "SHA256",
    "SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1203222200330322-1300312001022303-0100122323211131-2122203330333103-0300310121201012-2003000231312033-2322200333001333-1300020113311311"></a>

<a id="canonical-3212030113222212-1030031210112221-0332003311232311-3111132310322031-1011301330232230-2023201111313213-1321202332032103-3202101122223233"></a>

## selector property — values / 123231100331 / 7

Type: `"string"`. Optional.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("FullCertificate",
    "UseSubjectPublicKey"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002210202113330-3112021220110121-3332220111223312-2210200320000032-1220301110021013-2312022012013202-1110022011101030-3033013022200322"></a>

## Next pages — values / 123231100331 / 8

- [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-1110210111012102-2031202010320123-1132320121232332-2002023031230032-2022213032232233-1212213101223203-1103321022220000-0330112323013122)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3331231312030132-3010311201030021-0330032020003003-1111312303203000-3232301002223130-3103200221113020-1202122130031211-1301211330201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200300103031113-0122102002132021-1103032232121120-1012023311310323-3201022132300310-3023030021132001-2212223010013201-1222003133003223"></a>

## primary.default_rr_set_group.txt_record — txt_record / 031322103322 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.txt_record

<a id="canonical-1101322310230312-3222303222113203-2322030321021101-3321113220132023-2012323110330200-0103131230030110-3300121021102002-1022211223103031"></a>

Type: `"object"`. single nested block, Optional.

DNSTXTResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
txt_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222220101012001-0210111023320311-1313022322121022-1323322330102023-0102333201030203-1122001030112210-2303202032231101-3232120123220220"></a>

## Direct properties — txt_record / 031322103322 / 3

<a id="canonical-1021122013021120-1111100310231012-1003331323020003-1022023030313013-3203012320111313-0312211310312032-2000112201121211-3300232121200001"></a>

<a id="canonical-0112001201203210-0013233010121312-2103313032031132-2231101300102120-2223131000331222-3211203202333103-1302330212021303-2231102100023221"></a>

## name property — txt_record / 031322103322 / 4

Type: `"string"`. Optional.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3203321001320331-0232131310021333-3320331222001200-1010030112223331-2002103203121112-0011011311332130-1010001210210030-1012311123120023"></a>

<a id="canonical-0112110321010211-2003100131220303-3131131303210330-0221131310112120-2313310321013303-3212300013311202-0113023211020203-3310211122101232"></a>

## values property — txt_record / 031322103322 / 5

Type: `["list", "string"]`. Optional.

Text. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3101303130121212-1320331113322033-3022102002113323-2120010100100112-3230023321013201-3202032211111030-3330111310331302-3012330323110013"></a>

## Next pages — txt_record / 031322103322 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0001013031002232-2323220130321002-0110031123323021-1302303333101221-0212302022021323-3331121010231120-3131130311320003-3312131103003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311302102102011-3210000030310001-2120202231011322-0320320000122100-2100101031213303-3331130201012023-3233012303210011-1213220001320001"></a>

## primary.default_soa_parameters — default_soa_parameters / 301100112211 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.default_soa_parameters

<a id="canonical-3200211123121101-0222032221231303-2032000002003211-0010311331231033-3032330230003031-1133110003111120-2222022011330333-3102013131201222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default soa parameters.

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

Terraform syntax:

```terraform
default_soa_parameters = {}
```

<a id="canonical-3220032232301202-3230213332321211-1131200120001130-3221002101002223-1021200211003332-0221320020120311-0212322121223031-1213322201320212"></a>

## Direct properties — default_soa_parameters / 301100112211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333301032210020-0112003101103330-0022333332012202-3110122001123002-1020202332120330-3111300222023302-0030112011302311-3133002102230002"></a>

## Next pages — default_soa_parameters / 301100112211 / 4

- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301321130322033-2220232123213303-1231103221301133-3100211132111130-3211332333021103-2321303331110122-1102111133310010-2223001111100221"></a>

## primary.dnssec_mode — dnssec_mode / 203031230303 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.dnssec_mode

<a id="canonical-2010020000212012-0312302321031122-0031232110200121-2021212131320032-1030321233203212-1122212300013010-0302322103301010-2301132300330100"></a>

Type: `"object"`. single nested block, Optional.

DNSSEC Mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
dnssec_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103123030000331-2023000320303000-1113100211123321-0231133101212122-2322300131310230-2232311110132322-1133131303203303-0210113211033100"></a>

## Direct properties — dnssec_mode / 203031230303 / 3

- [disable_spec](resources--dns_zone--reference--group-002.md#canonical-2110111011300011-0110222200202000-2331001010300133-2101102011023011-3210111013013002-2001331000032223-1122311331130203-1030210321032303): complete subsection reference.

- [enable](resources--dns_zone--reference--group-002.md#canonical-1310313012332122-1030002133303121-1003102013010323-3332112323301223-3110312030122203-0312021133201032-1100323202001132-2131130301303301): complete subsection reference.

<a id="canonical-0122002000020331-1302002212222020-2102022113020221-0021222103013330-2133110030113101-3123333233301303-0102331111200113-3212130231110312"></a>

## Next pages — dnssec_mode / 203031230303 / 4

- [primary.dnssec_mode.disable_spec](resources--dns_zone--reference--group-002.md#canonical-2110111011300011-0110222200202000-2331001010300133-2101102011023011-3210111013013002-2001331000032223-1122311331130203-1030210321032303)
- [primary.dnssec_mode.enable](resources--dns_zone--reference--group-002.md#canonical-1310313012332122-1030002133303121-1003102013010323-3332112323301223-3110312030122203-0312021133201032-1100323202001132-2131130301303301)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2110111011300011-0110222200202000-2331001010300133-2101102011023011-3210111013013002-2001331000032223-1122311331130203-1030210321032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200303200232300-2200320022033022-2331230110311330-2023111301302221-3012210033101110-1033023300201301-1110123110210303-0323300200320011"></a>

## primary.dnssec_mode.disable_spec — disable_spec / 301213111230 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102)
- primary.dnssec_mode.disable_spec

<a id="canonical-0102011201220321-3223102212323133-1131213003012321-2300222021123323-2111313310131321-1033100302132121-0201310110333002-1330301122030221"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-1001231022200122-2033313312323030-0133322232300321-2112232022220101-2021212103211103-1131300002031310-2220132123300332-2022320032120332"></a>

## Direct properties — disable_spec / 301213111230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303111131111122-3313212322131103-0002120313101200-1000111132231203-1300001011330010-3302032013332210-2222221020111031-3100111320321303"></a>

## Next pages — disable_spec / 301213111230 / 4

- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1310313012332122-1030002133303121-1003102013010323-3332112323301223-3110312030122203-0312021133201032-1100323202001132-2131130301303301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330113212010221-0213030031200233-0213333201230302-2101312123030320-1301000221300212-3021133031020101-3013230023030332-2333122112230212"></a>

## primary.dnssec_mode.enable — enable / 102230320000 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102)
- primary.dnssec_mode.enable

<a id="canonical-2203011132121032-2323023131312322-1021021231331220-0131021102332033-3200331321131112-1130030220013101-2311023022113003-0200222031102223"></a>

Type: `["object", {}]`. Optional.

Enable. DNSSEC enable.

Upstream description:

DNSSEC enable.

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
enable = {}
```

<a id="canonical-0021211233221132-2132002233001130-2332020210130122-0333031331130303-2021312222231230-0210303122210301-2021210232112002-2301103013302123"></a>

## Direct properties — enable / 102230320000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022002313130202-2123022312100010-1200311033233130-0232120131030010-0322001301011111-1123311233203201-1110202231200010-0110033112212102"></a>

## Next pages — enable / 102230320000 / 4

- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030023201310033-0321023311211103-2202033122323221-3220020131231210-1101032100131021-3220110112333000-1022212130212101-2103103303331013"></a>

## primary.rr_set_group — rr_set_group / 102031023200 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.rr_set_group

<a id="canonical-2322001101312122-2303120132022113-0220120321000013-2022230220112311-0030333011013203-2232310232111211-1302303330103133-2221200303231123"></a>

Type: `"object"`. list nested block, Optional.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rr_set_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030212300132032-3102220032322332-2000101033201020-2203000033223122-1222023332230002-3131002233012123-1312131110301321-2021302212023013"></a>

## Direct properties — rr_set_group / 102031023200 / 3

- [metadata](resources--dns_zone--reference--group-002.md#canonical-0030212320120101-2321031322102202-0232302322102120-0030103101210333-2323021020120003-3300023212022220-2003202133231201-1301212100013303): complete subsection reference.

- [rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313): complete subsection reference.

<a id="canonical-3211301033011232-1333122233331121-0232120333031000-1203120130200233-0002003333131102-0000031121021002-0220210231122111-1001202300200330"></a>

## Next pages — rr_set_group / 102031023200 / 4

- [primary.rr_set_group.metadata](resources--dns_zone--reference--group-002.md#canonical-0030212320120101-2321031322102202-0232302322102120-0030103101210333-2323021020120003-3300023212022220-2003202133231201-1301212100013303)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0030212320120101-2321031322102202-0232302322102120-0030103101210333-2323021020120003-3300023212022220-2003202133231201-1301212100013303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301113111202230-1120310012130123-0131012132310031-1121301122113102-0233011331123300-0310001011331122-1321230000031330-3333000000333130"></a>

## primary.rr_set_group.metadata — metadata / 330310023033 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- primary.rr_set_group.metadata

<a id="canonical-1002010330023123-0133011030001021-0201203112331312-3203023122201120-3010100112102002-2211331230303303-3100201021201011-0301202201111300"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112023223321121-0122300231232023-0133221111303203-2133120313221130-2111230233313213-0323320231033320-0002313303101110-1130002212023003"></a>

## Direct properties — metadata / 330310023033 / 3

<a id="canonical-1322010030310310-0013023022313121-1321130201203213-1020303010211132-1333232101323021-3001232033002200-3021331312303202-0200220301133203"></a>

<a id="canonical-1303011112301201-3001310211102020-1302211122020200-2222332233022212-2212312303111303-1022000010202131-3010001203313311-3103331030312321"></a>

## description_spec property — metadata / 330310023033 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2013210032112020-2012212313210202-3310000230131101-3102233100030210-0333311232132132-3021313113233010-0110022333233002-2300213020220012"></a>

<a id="canonical-2201130001033030-3311333033231033-3002013203211133-0112302303022333-1310030021213311-1021213321322320-1200213102203000-3113323120320112"></a>

## name property — metadata / 330310023033 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3010323132201012-3113311011020202-0012010033020333-1322002321320022-0102100200230020-2223011301200001-1223002112220132-0220212120203222"></a>

## Next pages — metadata / 330310023033 / 6

- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110033313021321-3330300303320203-0133311002331003-1000311201130030-3132022212333232-0203322112323123-0202023331133032-3233022222212331"></a>

## primary.rr_set_group.rr_set — rr_set / 320320333202 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- primary.rr_set_group.rr_set

<a id="canonical-1321132121303313-2303003201012122-1132223021003222-3003122211220200-2032321003203202-2212100213021102-0323333113103201-2011103021121101"></a>

Type: `"object"`. list nested block, Optional.

Resource Record Sets. Collection of DNS resource record sets.

Upstream description:

Collection of DNS resource record sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ttl"),
  validators.ConflictingListObjectAttributes("a_record",
    "aaaa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("tlsa_record",
    "txt_record")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

Terraform syntax:

```terraform
rr_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300323211013030-2011320002002230-1002333121130213-1322113223023013-2033023212112221-2333313002222110-1201033132100102-1000321300002010"></a>

## Direct properties — rr_set / 320320333202 / 3

- [a_record](resources--dns_zone--reference--group-002.md#canonical-1120220211221011-3230322322033011-1233032203220321-1102111001122332-1332121133201301-1033322220200120-2313023233322321-3010230212113003): complete subsection reference.

- [aaaa_record](resources--dns_zone--reference--group-002.md#canonical-1333023102122010-3110332232301001-2120203012203220-3011102200230000-3223333103313110-1331230111222313-1203023033000321-0112021323323303): complete subsection reference.

- [afsdb_record](resources--dns_zone--reference--group-002.md#canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330): complete subsection reference.

- [alias_record](resources--dns_zone--reference--group-002.md#canonical-0233221001322012-3112103303232212-1312101110101110-3121331300013112-0011002002223132-1231003102222231-1212201121230333-3331020002012320): complete subsection reference.

- [caa_record](resources--dns_zone--reference--group-002.md#canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100): complete subsection reference.

- [cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133): complete subsection reference.

- [cert_record](resources--dns_zone--reference--group-003.md#canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101): complete subsection reference.

- [cname_record](resources--dns_zone--reference--group-003.md#canonical-1103130121131032-2002012232031213-3122130221012010-2331212231230031-3110031301310331-0321330030021320-3212001113023221-1011201311323322): complete subsection reference.

<a id="canonical-2133332032212210-0123213230222121-3123113232120122-1010332211233120-3100320331130120-0330130000220032-2113103331312030-1322231031232211"></a>

<a id="canonical-0011311101130102-2123210002031033-2031213001223031-1022112232002011-1221010200002211-2311131331032220-3322323302210001-3312032303101320"></a>

## description_spec property — rr_set / 320320333202 / 4

Type: `"string"`. Optional.

Comment. Human-readable description text

- [ds_record](resources--dns_zone--reference--group-003.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223): complete subsection reference.

- [eui48_record](resources--dns_zone--reference--group-003.md#canonical-3030020123311102-0330002300121201-0201032021303132-0331310302301121-2002232330201131-3201010010133002-1121220200310131-0121101001003201): complete subsection reference.

- [eui64_record](resources--dns_zone--reference--group-003.md#canonical-1123232112333312-0301023300332122-0100233233030001-1023002032221030-0013031220223001-2301310111033331-0120133033311203-0112320013203033): complete subsection reference.

- [lb_record](resources--dns_zone--reference--group-003.md#canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203): complete subsection reference.

- [loc_record](resources--dns_zone--reference--group-003.md#canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202): complete subsection reference.

- [mx_record](resources--dns_zone--reference--group-003.md#canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202): complete subsection reference.

- [naptr_record](resources--dns_zone--reference--group-003.md#canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223): complete subsection reference.

- [ns_record](resources--dns_zone--reference--group-003.md#canonical-3112102233211200-0313123203112300-1130013102230113-2011023232120220-1302012230210220-3202131201133110-0011021300201313-1102311332102132): complete subsection reference.

- [ptr_record](resources--dns_zone--reference--group-003.md#canonical-3232000310031110-1111032130320130-1330010220023223-2023200213300002-3123322311111130-2121011102311023-1203013003000322-3300000321210110): complete subsection reference.

- [srv_record](resources--dns_zone--reference--group-003.md#canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121): complete subsection reference.

- [sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311): complete subsection reference.

- [tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110): complete subsection reference.

<a id="canonical-2212321033023303-0323001231312003-3210032310003003-0201001200102131-0131103232010111-2122312100213023-1303001300022011-0200102110001300"></a>

<a id="canonical-1110011133332213-1022030323123333-2122310130033021-2022221312003102-2223320232131332-2022222020201200-1010133303010000-2122330311220233"></a>

## TTL property — rr_set / 320320333202 / 5

Type: `"number"`. Optional.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](resources--dns_zone--reference--group-003.md#canonical-1121131111220220-0333230311100110-3220333331113330-3330020320131230-1033022131302013-0333232321103112-3113013123231300-0230232233102113): complete subsection reference.

<a id="canonical-0212002311120121-2022023122030113-3101312133323110-3121100020331022-0112022012032202-0011302032032033-1121003132200222-2211120022022131"></a>

## Next pages — rr_set / 320320333202 / 6

- [primary.rr_set_group.rr_set.a_record](resources--dns_zone--reference--group-002.md#canonical-1120220211221011-3230322322033011-1233032203220321-1102111001122332-1332121133201301-1033322220200120-2313023233322321-3010230212113003)
- [primary.rr_set_group.rr_set.aaaa_record](resources--dns_zone--reference--group-002.md#canonical-1333023102122010-3110332232301001-2120203012203220-3011102200230000-3223333103313110-1331230111222313-1203023033000321-0112021323323303)
- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330)
- [primary.rr_set_group.rr_set.alias_record](resources--dns_zone--reference--group-002.md#canonical-0233221001322012-3112103303232212-1312101110101110-3121331300013112-0011002002223132-1231003102222231-1212201121230333-3331020002012320)
- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101)
- [primary.rr_set_group.rr_set.cname_record](resources--dns_zone--reference--group-003.md#canonical-1103130121131032-2002012232031213-3122130221012010-2331212231230031-3110031301310331-0321330030021320-3212001113023221-1011201311323322)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [primary.rr_set_group.rr_set.eui48_record](resources--dns_zone--reference--group-003.md#canonical-3030020123311102-0330002300121201-0201032021303132-0331310302301121-2002232330201131-3201010010133002-1121220200310131-0121101001003201)
- [primary.rr_set_group.rr_set.eui64_record](resources--dns_zone--reference--group-003.md#canonical-1123232112333312-0301023300332122-0100233233030001-1023002032221030-0013031220223001-2301310111033331-0120133033311203-0112320013203033)
- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203)
- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202)
- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202)
- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223)
- [primary.rr_set_group.rr_set.ns_record](resources--dns_zone--reference--group-003.md#canonical-3112102233211200-0313123203112300-1130013102230113-2011023232120220-1302012230210220-3202131201133110-0011021300201313-1102311332102132)
- [primary.rr_set_group.rr_set.ptr_record](resources--dns_zone--reference--group-003.md#canonical-3232000310031110-1111032130320130-1330010220023223-2023200213300002-3123322311111130-2121011102311023-1203013003000322-3300000321210110)
- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110)
- [primary.rr_set_group.rr_set.txt_record](resources--dns_zone--reference--group-003.md#canonical-1121131111220220-0333230311100110-3220333331113330-3330020320131230-1033022131302013-0333232321103112-3113013123231300-0230232233102113)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1120220211221011-3230322322033011-1233032203220321-1102111001122332-1332121133201301-1033322220200120-2313023233322321-3010230212113003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032100101230123-1113210131021132-3022010123232223-2101002320202122-1123112233202222-1031331022100011-3033301133112311-3101312302021003"></a>

## primary.rr_set_group.rr_set.a_record — a_record / 230132331101 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.a_record

<a id="canonical-2031310110000111-1322322130201330-3323301001110212-3133310330133013-0210013033102131-2203230200200030-2213013220132213-1301102103210031"></a>

Type: `"object"`. single nested block, Optional.

DNSAResourceRecord. A Records

Upstream description:

A Records

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
a_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231230332033100-0213130330312302-1200032101032022-3030022122131311-0121123210132122-2303102332120333-3130332000200120-0220101020101203"></a>

## Direct properties — a_record / 230132331101 / 3

<a id="canonical-0220023032013002-1221230021322120-1213110113102212-2212322313110302-0021203310120013-0333222231333213-3301220311113111-3020010330101031"></a>

<a id="canonical-1002100230330110-2103310230022032-3332013223103012-0121012112322302-1121000023303133-3223303222133012-2023310310021132-0032023223230303"></a>

## name property — a_record / 230132331101 / 4

Type: `"string"`. Optional.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0300022002231102-2302030030032203-3030233303233321-0223032331312033-2323030002001021-3310200220101300-3022110323013323-0333223132133122"></a>

<a id="canonical-0122120001312301-1231321130103132-0023030033331010-2300103100310330-2011221030113023-0002232333203113-2231210020003212-2122010233022101"></a>

## values property — a_record / 230132331101 / 5

Type: `["list", "string"]`. Optional.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131022031312323-1333200012202100-0111033301233003-0232001222200320-0121223301233203-1311201031120210-2020012003313312-1132231002210121"></a>

## Next pages — a_record / 230132331101 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1333023102122010-3110332232301001-2120203012203220-3011102200230000-3223333103313110-1331230111222313-1203023033000321-0112021323323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210322033030330-1213002110230331-0023002132001303-1032113010231000-1012001221331111-2032000330300200-0332003131013000-1001123110321330"></a>

## primary.rr_set_group.rr_set.aaaa_record — aaaa_record / 333303200102 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.aaaa_record

<a id="canonical-1323010102202331-0031123232222202-0211213021200023-3012223001120200-1132111003100012-2302000011113313-2120122120121020-0333031000330003"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
aaaa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312203323033111-3101303233223321-3211202313003332-0113200131303130-3000333030203130-1110313131220003-2131003201211303-3211113211033023"></a>

## Direct properties — aaaa_record / 333303200102 / 3

<a id="canonical-3223312000221212-3003122023110122-2003002023300333-0211011333112000-0012300333010312-2222110201231331-0203200221131130-0132121232031223"></a>

<a id="canonical-2001230121200100-2013132212010311-1220111021113022-1321133203100203-3321021300122102-3033022121300123-0021122130033333-3031310303020020"></a>

## name property — aaaa_record / 333303200102 / 4

Type: `"string"`. Optional.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3020221301220010-1032211021132313-1123220110010130-0122113131112121-2231112232001330-3103313032100111-2232111032011301-1022233132013200"></a>

<a id="canonical-0303133112311102-0132332003312203-2002231131003211-3023323211333311-2123011032032133-3331002223032133-1122213322113200-3202222322312213"></a>

## values property — aaaa_record / 333303200102 / 5

Type: `["list", "string"]`. Optional.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1012302010000021-3023220123021100-1333031320120211-3200023322232122-3311323002211033-1213322120103223-3211111213131133-2310012021210322"></a>

## Next pages — aaaa_record / 333303200102 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123133312121111-2303330101313103-3222133322112100-2122131033131323-2302310120222122-0113130010222211-1313103021302003-3210010000031313"></a>

## primary.rr_set_group.rr_set.afsdb_record — afsdb_record / 023203331303 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.afsdb_record

<a id="canonical-1212310322223223-2312001020221020-2230222312010033-2320000220311213-0003133123122130-2310303301101003-3111300121022223-2213303111133022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
afsdb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100311001010213-3001222011020132-2022212302102012-2102220332331232-0303301310010311-2320003020212132-3323112001312222-3120222332033311"></a>

## Direct properties — afsdb_record / 023203331303 / 3

<a id="canonical-1213302211210031-2133320200203000-0111102230320010-1221332300120121-0223003031321222-0303132322330003-0030220030203103-1112223023220031"></a>

<a id="canonical-2131213120320032-2002101012223030-2112023020030311-3120301333231121-2301330033232333-3101302120331110-3120132330201021-2322311303130031"></a>

## name property — afsdb_record / 023203331303 / 4

Type: `"string"`. Optional.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-3301223013203213-3100232023203123-0323311031321001-0003112110023320-2303310020231103-2322133102323010-1130031222122123-0110302230131311): complete subsection reference.

<a id="canonical-2212222132110312-1102322123033022-0111021322231211-3300010303311201-1222212212020000-1110020010001122-3311300230302332-2312321300320321"></a>

## Next pages — afsdb_record / 023203331303 / 5

- [primary.rr_set_group.rr_set.afsdb_record.values](resources--dns_zone--reference--group-002.md#canonical-3301223013203213-3100232023203123-0323311031321001-0003112110023320-2303310020231103-2322133102323010-1130031222122123-0110302230131311)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3301223013203213-3100232023203123-0323311031321001-0003112110023320-2303310020231103-2322133102323010-1130031222122123-0110302230131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320223310331303-0310102023310002-0323130202110321-3233300312220333-0110032231100032-3201103311020102-3311213321122223-2210133032213331"></a>

## primary.rr_set_group.rr_set.afsdb_record.values — values / 011131102020 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330)
- primary.rr_set_group.rr_set.afsdb_record.values

<a id="canonical-2302200222330321-1300110110323123-1303233000130322-3201222130200231-0003332331333121-3201222231130021-0232113000202130-0232132312031203"></a>

Type: `"object"`. list nested block, Optional.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("hostname")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200013021210211-0322301020020131-0003132233212010-1121211010210310-2021130212013001-3130311101333232-1222020202231311-0130100023200102"></a>

## Direct properties — values / 011131102020 / 3

<a id="canonical-2102002201003330-2002101221012223-3001203121312002-3211211121013330-3111312122230022-1210303123021302-3111023333013132-2313221302033331"></a>

<a id="canonical-2302233221221323-2132133300001323-0000001033011332-1020002330231321-3211330122103210-0320301323230331-2312233210313102-3331002322101102"></a>

## hostname property — values / 011131102020 / 4

Type: `"string"`. Optional.

Server name of the AFS cell database server or the DCE name server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2000111133310013-2233120201210110-2220002303030010-2110000110303223-3010023113213023-3123100010310020-1201230020100323-3032301122200321"></a>

<a id="canonical-3021320123112210-0220313202002203-2120223321120113-3133121330203120-0120312301221001-3323021011133112-1131332223230322-2230320200030333"></a>

## subtype property — values / 011131102020 / 5

Type: `"string"`. Optional.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001210312132310-2032301312121221-1321323133301213-3033122221012232-3013300301130330-1211033233123221-2003022233121301-0123231310113112"></a>

## Next pages — values / 011131102020 / 6

- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0233221001322012-3112103303232212-1312101110101110-3121331300013112-0011002002223132-1231003102222231-1212201121230333-3331020002012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303310201112310-3110332010032330-2133231202030131-3233233133302131-2210102233300311-0003313303021022-2333320200032101-3322210003100330"></a>

## primary.rr_set_group.rr_set.alias_record — alias_record / 320300110232 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.alias_record

<a id="canonical-3023031303112002-1011220313330031-3300012333020203-1103120213122123-0332110100122323-0321332213233010-1003301332021321-3322231133202100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for alias record.

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
alias_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131333332102010-0110322331332032-0322120002013300-1201312021221303-0210033322202310-0133120131233333-0003023313130211-0300122100332021"></a>

## Direct properties — alias_record / 320300110232 / 3

<a id="canonical-3203103303223000-2221133301210312-1313103330113003-1000031030201221-0122121222313102-1221020030310013-3021011011232033-3231211111121212"></a>

<a id="canonical-0010110002000123-2302210202311120-1212132320233231-1232320030312300-1232133321121121-0213232122102000-0133303220133301-3213231103033211"></a>

## value property — alias_record / 320300110232 / 4

Type: `"string"`. Optional.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1121132310310210-2103302221221202-2312211203001310-2200113330020033-0020323303132203-2211302103022121-1131111122333302-3130022100231332"></a>

## Next pages — alias_record / 320300110232 / 5

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023203333202013-1011032320303301-0201232002233202-1232221021301002-2321212310231211-1132310331003322-2202331203312022-2300221332220120"></a>

## primary.rr_set_group.rr_set.caa_record — caa_record / 031220233003 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.caa_record

<a id="canonical-3203022223320231-1032100010021023-3102132232202011-3221031132212002-2332032001000210-3210130210033031-0230230202331030-0132232201123103"></a>

Type: `"object"`. single nested block, Optional.

DNSCAAResourceRecord.

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
caa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003121110322311-1333020002311220-1123101210222133-1212102231013001-3021320030101103-3300211000112021-0101103230130123-3201021100101030"></a>

## Direct properties — caa_record / 031220233003 / 3

<a id="canonical-0202232310230231-0100333010221323-2311013202202210-1230220200223232-0130303212131312-1131222311113120-1033123120003232-3010203312311203"></a>

<a id="canonical-0113313000210133-3000121002222023-2011200030112321-3011320330000112-3132201202320101-3122133211322311-0323002313131223-2301301012021013"></a>

## name property — caa_record / 031220233003 / 4

Type: `"string"`. Optional.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-3233301110320230-2121202201132103-3002310301013122-2200203211110111-0012201220000310-0102313030033200-1131210233232132-0110201023022101): complete subsection reference.

<a id="canonical-1321120213320121-1211030101031011-3302013220021130-2023031033212010-2303120121020112-1302010231101212-1113230323022100-1223030213113321"></a>

## Next pages — caa_record / 031220233003 / 5

- [primary.rr_set_group.rr_set.caa_record.values](resources--dns_zone--reference--group-002.md#canonical-3233301110320230-2121202201132103-3002310301013122-2200203211110111-0012201220000310-0102313030033200-1131210233232132-0110201023022101)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3233301110320230-2121202201132103-3002310301013122-2200203211110111-0012201220000310-0102313030033200-1131210233232132-0110201023022101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201202000202320-0321002213121023-2211333321100303-2313001022133123-2021301033111221-2101100313303001-3022021111222022-0202202121122020"></a>

## primary.rr_set_group.rr_set.caa_record.values — values / 102113111012 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100)
- primary.rr_set_group.rr_set.caa_record.values

<a id="canonical-2113020310310122-0010131022220111-2010230133103022-0213103110113121-1002000132023232-2111320322131310-0112002102320313-1102302232201030"></a>

Type: `"object"`. list nested block, Optional.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222211110312010-2300032322200311-2302332102010112-0112213110323012-3221211120321021-3111100103331310-3021312233302122-0031003022020312"></a>

## Direct properties — values / 102113111012 / 3

<a id="canonical-0301013201232232-3231012001032130-2102210023001223-3113003232330333-1230031010002122-3123133202333113-0212022022112022-2103100200100311"></a>

<a id="canonical-0331302203333120-2223302031012212-0312000131213211-1001121032000313-3011010200133113-3321121203311021-0000122200103200-1012001301310133"></a>

## flags property — values / 102113111012 / 4

Type: `"number"`. Optional.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-0330230220122010-2001122300032033-0202022222323200-0233331130103322-1131320300310111-1332321010302331-0203103301013320-3231111120230011"></a>

<a id="canonical-3100301101320030-2321221210323102-0202103322232211-1020212003300030-1323102232031030-3223132332213011-2301012301322311-1220001233113323"></a>

## tag property — values / 102113111012 / 5

Type: `"string"`. Optional.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("issue",
    "issuewild",
    "iodef"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-0000131020323322-3312101311333222-3310002101231232-2213233023311213-3002212311223300-0310121023110323-1301230212012130-0211303130103132"></a>

<a id="canonical-0113112323002323-2323113213212003-3011102222202020-0120010203111133-3202213332131231-3312122323323300-0133322331201101-3001002223031102"></a>

## value property — values / 102113111012 / 6

Type: `"string"`. Optional.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2122331310211032-0130302130020322-0232020111320220-2013001220120112-0112213103220211-0220103312320203-1210102023201022-0003110113000012"></a>

## Next pages — values / 102113111012 / 7

- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230200010210003-2223301210223102-1103002303320013-2120023000212111-0321233110112030-2313230233232022-1200231211230102-3022001212231321"></a>

## primary.rr_set_group.rr_set.cds_record — cds_record / 320300332131 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.cds_record

<a id="canonical-3012220321222313-2102100303223132-3320001303231312-1113220202122233-3233033202202311-3303302130320322-3130001301210120-3230110112300123"></a>

Type: `"object"`. single nested block, Optional.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
cds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110121023220131-0112103103111233-1103110030110032-2113210103113233-0232130033131322-2310320112313032-0033313120130032-3313000322222031"></a>

## Direct properties — cds_record / 320300332131 / 3

<a id="canonical-1202020213032001-1221001000120313-0200233210221010-1120001323311232-0231212323131302-1220021102102332-1111103201120221-0023103032201132"></a>

<a id="canonical-0122321102113013-1123212013311321-3221302311033230-3303303012322231-0221320012122012-2231301311000023-2130221300233202-1231012330021311"></a>

## name property — cds_record / 320300332131 / 4

Type: `"string"`. Optional.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033): complete subsection reference.

<a id="canonical-2021121311011300-2131331220222120-1122213023102211-2333211030203321-2223003223000032-0232121200330020-3130020322301300-3330233330002101"></a>

## Next pages — cds_record / 320300332131 / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
