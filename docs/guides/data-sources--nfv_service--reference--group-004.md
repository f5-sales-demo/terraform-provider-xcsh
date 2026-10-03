---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-1011213113222321-0200333200212001-1202112233013101-3132101332211123-3331203310332222-3013213302321320-0130221301222223-3113332130031231"></a>

## Direct properties — private_key / 011213022110 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-2002003200130322-3223321312321232-3113023310323022-2012131022123111-0022311101201101-2312033012223223-1213302232112300-0201320233032120): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-1102110202003220-1120303130223002-1021013031233113-1132102330322333-0032331121300030-3030020011000300-0033301323103302-1223110212033221): complete subsection reference.

<a id="canonical-0011012011010312-0212013013303103-3103110110332121-2013021200003212-1330003203130330-1020103030303222-1222030133012301-0132011202321233"></a>

## Next pages — private_key / 011213022110 / 4

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-2002003200130322-3223321312321232-3113023310323022-2012131022123111-0022311101201101-2312033012223223-1213302232112300-0201320233032120)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-1102110202003220-1120303130223002-1021013031233113-1132102330322333-0032331121300030-3030020011000300-0033301323103302-1223110212033221)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2002003200130322-3223321312321232-3113023310323022-2012131022123111-0022311101201101-2312033012223223-1213302232112300-0201320233032120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322133302100130-3001323231102232-0020023210210101-3102301311230132-0213130100100211-1222002220131230-2032310130202023-1222322231021312"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info — blindfold_secret_info / 212200001200 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1331100032030202-1322113010113230-3223011120122320-1323023030212123-3111102200332001-0212213310201023-3131213223212011-3202310100323131)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info

<a id="canonical-0131101220320321-0313303200021110-0313202000312312-2213031323010011-3111202320133003-3100022210003103-3110230231122210-1012203221031221"></a>

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

<a id="canonical-3300101211010123-1232012002302303-1111323123132030-0001303110102030-3032233101221302-3100132332331033-1312222313011012-1330113310022223"></a>

## Direct properties — blindfold_secret_info / 212200001200 / 3

<a id="canonical-2333321302102031-3012222023322113-2130333010020013-2313320020102010-3221033020102012-1131301303013121-0231331120012110-3331200012112102"></a>

<a id="canonical-0122321010110301-3303010322200100-0202331232222331-2110310023130321-0232010310132320-1200220122012011-0020301120130102-1313212300121333"></a>

## decryption_provider property — blindfold_secret_info / 212200001200 / 4

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

<a id="canonical-1313233132303031-2332331330330013-3332020103112003-2033223300122100-1103212221321321-3310303201031110-1201012133231122-2030113200321220"></a>

<a id="canonical-3011000200003122-1230130110132103-2203330121330210-0131103022001213-2223302031023231-2020121012103221-1233110212121112-2221003222322130"></a>

## location property — blindfold_secret_info / 212200001200 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2233130003103312-2300113023130212-1023220323031121-1211003201132320-3320112313222013-2311013001103123-1302133302203310-1321301223230233"></a>

<a id="canonical-0301323010023213-1210113333310112-0312122230231313-2212133311330320-2010023101233133-0010130122302322-3021300030020331-0301001310001333"></a>

## store_provider property — blindfold_secret_info / 212200001200 / 6

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

<a id="canonical-1203301123303300-1130133132001123-2230303000331212-1330201121010112-0022113323222312-1202103322301322-1231210203311130-3133230232023021"></a>

## Next pages — blindfold_secret_info / 212200001200 / 7

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1331100032030202-1322113010113230-3223011120122320-1323023030212123-3111102200332001-0212213310201023-3131213223212011-3202310100323131)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1102110202003220-1120303130223002-1021013031233113-1132102330322333-0032331121300030-3030020011000300-0033301323103302-1223110212033221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212012320200012-2032303010300133-0010320332131030-3102203033302022-0020333132331332-0330210201231010-0331332330231132-3032133001333011"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info — clear_secret_info / 112313202311 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-0210230323010131-3011230220332132-1312310010020002-2332233101111330-1202323330002030-1233232002132011-2130031030101332-3000013212030311)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-0320223322111202-0002212331032222-3013132301312002-1233000031330133-2010221021112300-0230312213113101-1031002322112211-1000303020321121)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1331100032030202-1322113010113230-3223011120122320-1323023030212123-3111102200332001-0212213310201023-3131213223212011-3202310100323131)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info

<a id="canonical-1303113033003331-1300002003301322-2300033220333211-3012213221302010-3202321231000330-3221223012312031-2033331000212301-2300113321320110"></a>

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

<a id="canonical-3013232023110330-2121013033211012-1113212223003011-3002120202330130-1232102320132231-1123132103331233-3002301003032323-0102110202030021"></a>

## Direct properties — clear_secret_info / 112313202311 / 3

<a id="canonical-0321131021223332-3113212120310012-2300210232031202-3133001331232303-0212122122200101-0032322122000233-3002132333321023-2103130130032321"></a>

<a id="canonical-1211321222230101-2332232223120132-2213130012030122-3322322022010012-0330300130130023-2210302113030012-2220230030103211-3023010032112011"></a>

## provider_ref property — clear_secret_info / 112313202311 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2033003333221223-2023033030102030-2313030011320331-3123112010221232-3010113333131013-3120221031322122-0330030110112112-3120211211312233"></a>

<a id="canonical-0131323020323321-1301210322212103-2302232012003122-1233222123312222-2023113321012011-1330303333111000-0212111302130333-2112330132332122"></a>

## URL property — clear_secret_info / 112313202311 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3122232312131121-0320203100033211-0323003211310113-1001203110333331-1131310032300321-0230111321003223-2113213010110231-2130033210222321"></a>

## Next pages — clear_secret_info / 112313202311 / 6

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1331100032030202-1322113010113230-3223011120122320-1323023030212123-3111102200332001-0212213310201023-3131213223212011-3202310100323131)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3033333023320323-0001020103321300-0132300322323200-3332031131310232-2330121300210232-2000130003110030-3321123030223130-0200133200122303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222002001020022-2223320202310201-3300002120010013-2322132231112023-1001322111321113-1303300033200032-0131001011123303-3203213202121120"></a>

## palo_alto_fw_service.aws_tgw_site — aws_tgw_site / 231232013200 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- palo_alto_fw_service.aws_tgw_site

<a id="canonical-1002223122113121-3102330210031330-2023303032233111-3110002013231221-0222001203101033-1103102012010210-1033311010133002-2320320332000021"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2302121123321323-2211331011030331-3020021310210130-3110232103200230-0030201010021202-3202313031322002-1330100111121311-2202021302012300"></a>

## Direct properties — aws_tgw_site / 231232013200 / 3

<a id="canonical-2011200112100303-2113221130301032-1020031310220113-0312012303011002-2313023131233031-2203312220133303-2101132020311130-3310033321003230"></a>

<a id="canonical-1120033232131313-2001312021123202-2131003003200112-2212212230221100-0200122213001003-1001000132312313-1203331020023200-0003101120021203"></a>

## name property — aws_tgw_site / 231232013200 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3233201302011022-1103301222013221-0212130131033132-2230321233132313-3112033131102103-3021233222313212-3322200111113130-0330202012232210"></a>

<a id="canonical-1301233321233212-0230113231211021-1123332133101023-1211000132322010-0210313120113333-3121110302203103-2213031200332100-3331213110113332"></a>

## namespace property — aws_tgw_site / 231232013200 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1323010012311002-0321020310020020-2033310013012130-1302121232102033-1201031122210122-2022330033121232-0332330333032000-2030033200301222"></a>

<a id="canonical-0302301231030100-3112030222230210-3222203212133033-1133331022032232-3301311313000312-3323122221310101-3313303113313132-0133021332032301"></a>

## tenant property — aws_tgw_site / 231232013200 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3132133032303030-2312013002232232-1012232212210013-1023100223310122-3103232222310233-3001330231031111-1311010311303311-3210300332030223"></a>

## Next pages — aws_tgw_site / 231232013200 / 7

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1203222020023202-0332023100202302-3112302323202100-2031120110113020-2003012123121220-2321212122231220-3201112333213233-1120321103332100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311032132302321-2310010301220323-0033311311322120-2322003311310103-2021120220331313-3101101212203320-0311102002222221-2111333220210333"></a>

## palo_alto_fw_service.disable_panaroma — disable_panaroma / 020122231300 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- palo_alto_fw_service.disable_panaroma

<a id="canonical-3221023100000033-1202313210103312-3322133323120231-3000201322320020-3021210032113302-1210030220232012-0013231312030123-1130332311001022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable panaroma.

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

<a id="canonical-0301222103121331-2003213200120221-1220333310303300-3032103023302211-3212102020113031-3310231333133330-1131022110231211-0021123212212110"></a>

## Direct properties — disable_panaroma / 020122231300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000030110000332-2011313111330001-3311010303102322-3203300223330023-3323013231232223-2032031031001203-3233100002301202-3033121113331120"></a>

## Next pages — disable_panaroma / 020122231300 / 4

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2110223121000233-3221320303312122-2000110302020133-0131111331123033-3303013000311223-3302120201222211-2220203233302110-3213113221022320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101221200231132-1012312012031123-3123110112021223-3200112202032232-2130211033002011-2301332332201001-2321110311010010-0102112022132322"></a>

## palo_alto_fw_service.pan_ami_bundle1 — pan_ami_bundle1 / 333213212301 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- palo_alto_fw_service.pan_ami_bundle1

<a id="canonical-0330002021030120-2230211133322211-0322022233201113-0121123002222100-2103200023222111-0023201000002102-1333333220100033-3012321122230133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pan ami bundle1.

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

<a id="canonical-0132223330013131-0133223310311123-0212011303101221-2200101100300013-0121213302302200-1020102333032232-1113133033230232-0211322323330222"></a>

## Direct properties — pan_ami_bundle1 / 333213212301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120310313113302-3211313310331100-0331311121221220-2132023012303200-3000122120231223-3220312021100102-2222203023223102-1221312130303210"></a>

## Next pages — pan_ami_bundle1 / 333213212301 / 4

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0123023022211213-3102321300301221-2221111210200132-2001212120303313-0301133303031123-2030312313201113-0210201002303230-0302312003233303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300301202220013-2322202210321321-2322213131330213-2300003223012232-0221032222320210-0310133111300202-3223120300312011-3022330032103122"></a>

## palo_alto_fw_service.pan_ami_bundle2 — pan_ami_bundle2 / 013301333212 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- palo_alto_fw_service.pan_ami_bundle2

<a id="canonical-2223112200023231-1001103221120030-2303002012232110-3011010221231033-3213030210211200-1320232030102222-0001210032111120-0120232002201202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pan ami bundle2.

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

<a id="canonical-1112300202113330-0013202333123213-1331201202103102-1303110100003211-0131300111333232-1011011002030220-1201202032323020-2102323011102102"></a>

## Direct properties — pan_ami_bundle2 / 013301333212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312110032112120-3123011003301110-3132322210133123-2123323301122220-0313310323032331-0133111023111120-0332122120113200-0302312211212100"></a>

## Next pages — pan_ami_bundle2 / 013301333212 / 4

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011000221312321-1221233122232131-0001001103111022-3230101120102133-2113320333020020-2223121312010233-0120021131001131-3013131332333033"></a>

## palo_alto_fw_service.panorama_server — panorama_server / 030211203011 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- palo_alto_fw_service.panorama_server

<a id="canonical-3312312313220311-0222321000231020-1113120020312122-0023033300023311-3103313121032222-3120303003203103-2133330320302300-1300212000230212"></a>

Type: `"single"`. Computed.

Configuration parameter for panorama server.

Upstream description:

Panorama Server Type.

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

<a id="canonical-3312122203011132-1302311303123020-1021210200001230-3302000213122212-1212312332020331-0020302211003203-0320213201210230-0112001312210232"></a>

## Direct properties — panorama_server / 030211203011 / 3

- [authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121): complete subsection reference.

<a id="canonical-3200133030102030-0211013233110300-1122121200100202-0222210221223101-1232310232212320-0023200231323303-0322332012323231-1313312001212231"></a>

<a id="canonical-3320113122123303-0130111222321132-3302033301323021-3121030203133213-2233130110311232-3230213333210133-2320013033111000-0203310222120330"></a>

## device_group_name property — panorama_server / 030211203011 / 4

Type: `"string"`. Computed.

Device Group Name. Device Group Name.

Upstream description:

Device Group Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3223212320221030-1333020123302023-0101101001003301-3211021212320331-0033113223322211-0310130210013322-3302220101220310-0031120212023231"></a>

<a id="canonical-1302203003120120-2112000100010032-0233021233313132-1130333110212313-2312130103332112-3131230200113200-1233100222332102-0132310233313100"></a>

## server property — panorama_server / 030211203011 / 5

Type: `"string"`. Computed.

Panorama Server Address to which the firewall should connect to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3323031213320330-1120211030313102-0001202233313203-0202020310312220-2030003003132032-3033200212320223-0220332201231320-0101122231113330"></a>

<a id="canonical-3323310112002010-0033121130313320-2001220200123321-3311223310022033-3222012133001322-1210213313110223-0233002233102230-0310333232302012"></a>

## template_stack_name property — panorama_server / 030211203011 / 6

Type: `"string"`. Computed.

Template stack name. Template Stack Name.

Upstream description:

Template Stack Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1131201310221012-0212233213133210-3011130033012132-0221022000221032-3012231130011303-2003221323231101-3313233022230102-2120030101102112"></a>

## Next pages — panorama_server / 030211203011 / 7

- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101030012001031-0011133331132123-1332201101312102-1112311202212230-0102311132332300-0311220010303200-2333030113332132-0312100310033123"></a>

## palo_alto_fw_service.panorama_server.authorization_key — authorization_key / 031023133221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200)
- palo_alto_fw_service.panorama_server.authorization_key

<a id="canonical-1300201323132212-0130320020131010-0012100122031033-2021312203100320-2011210223030021-2111020223323311-1100100120202203-1013321001200303"></a>

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

<a id="canonical-1100312132320203-3212302033221111-3100200111110021-3223113201000301-2112202203100210-2110013002233303-3301103000023100-3222032031302233"></a>

## Direct properties — authorization_key / 031023133221 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-1121130120121310-0331233013221333-0003010022002013-3101111012323101-2121030111003010-2331222201130332-0012120320302330-2231103003323123): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-0233100123303130-3330212002130103-2030331001211012-0321203213130000-0022111123230020-0132330001123132-0201021202212220-0200321303231003): complete subsection reference.

<a id="canonical-3002010033302033-1113120220302332-0011112212012110-0011310012301223-0221010221322332-3121232012000113-2202202013231321-3221011013311303"></a>

## Next pages — authorization_key / 031023133221 / 4

- [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-1121130120121310-0331233013221333-0003010022002013-3101111012323101-2121030111003010-2331222201130332-0012120320302330-2231103003323123)
- [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-0233100123303130-3330212002130103-2030331001211012-0321203213130000-0022111123230020-0132330001123132-0201021202212220-0200321303231003)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1121130120121310-0331233013221333-0003010022002013-3101111012323101-2121030111003010-2331222201130332-0012120320302330-2231103003323123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120130033322113-3001231123131200-1013320113002023-0122130301100222-2101230220223000-3022313023321020-2201123313121010-2322213220311203"></a>

## palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info — blindfold_secret_info / 031323113221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200)
- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121)
- palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info

<a id="canonical-3032113131223123-2112301032200002-3321232112022210-1212230031332322-2023203332302323-3323121303203323-2222303223130133-2323112211012310"></a>

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

<a id="canonical-0110111223112312-2130203020303133-3231031232113230-1332102133220132-2331320300123103-0000112203032220-1131333333032301-2030011320231133"></a>

## Direct properties — blindfold_secret_info / 031323113221 / 3

<a id="canonical-3211002012201312-2102133313332013-0222032101022011-0103230132020210-2233122331330103-1223320022210030-3032212332212131-0303122321120231"></a>

<a id="canonical-0001211220131112-1232322201222331-3131013301300130-3101333111032300-1333221313132030-3200113213203203-0023021030303330-0123321001321322"></a>

## decryption_provider property — blindfold_secret_info / 031323113221 / 4

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

<a id="canonical-1312020013310023-2033013132221001-1202201120000031-1132131100330111-0133300321000311-3310203231220333-3320031322121232-1000013323303012"></a>

<a id="canonical-0022110333210311-1212123220003211-0033223220302132-2301032322311033-1130021203230020-2202320023331123-1102132103131201-1323033020132300"></a>

## location property — blindfold_secret_info / 031323113221 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3330132002022100-0232331211122021-0030110101011202-2102031202010103-0023013221122303-3213322121110030-1000212210103221-1002203210202002"></a>

<a id="canonical-2313003133022031-3231313103132221-3022203103101312-0003220303202301-3012030130022201-2101113031100303-1332000023211333-3010232303011120"></a>

## store_provider property — blindfold_secret_info / 031323113221 / 6

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

<a id="canonical-2312021013310110-1021302300202322-1101112220032220-1010223331230221-3122023030310220-2203130013113300-0322311332113110-0331333011033331"></a>

## Next pages — blindfold_secret_info / 031323113221 / 7

- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0233100123303130-3330212002130103-2030331001211012-0321203213130000-0022111123230020-0132330001123132-0201021202212220-0200321303231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011302233312312-0002203313022221-0030002332100123-3010232022022212-0023213133232300-1131000320201003-3201210032313132-2130323013032123"></a>

## palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info — clear_secret_info / 212012113212 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200)
- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121)
- palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info

<a id="canonical-3331302012321122-2021321300322212-0320221013133233-1112323231110223-2113032232003300-3231031320202230-0313330031130123-2111021313003230"></a>

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

<a id="canonical-1011011003222012-2210210111030102-3102233211310203-0230003202021113-3223132300310233-0300020233022121-1031312322112113-0103230020332312"></a>

## Direct properties — clear_secret_info / 212012113212 / 3

<a id="canonical-2320201313233220-3130320020233012-2133121123222132-3132032000123233-0002023102103002-3010120110123013-3030203112033332-2202001132123231"></a>

<a id="canonical-1310103220321022-2332202313130322-3301003323313013-2200203100121300-2022213020013320-2200200112022110-0321101013100310-3131123202120003"></a>

## provider_ref property — clear_secret_info / 212012113212 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3033102111332302-0322302000133202-3230232331101200-1011213021300030-2130231330010010-3201023132001221-2033313002011120-2003201033023213"></a>

<a id="canonical-2302032102100302-3033333023230233-2331233202220001-2103330332022220-1012333020202122-0133130300131111-3123121331011013-3231001300103012"></a>

## URL property — clear_secret_info / 212012113212 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2333202210010301-3132231001221202-0020210003012001-0011003313221201-3131222212120202-2033313331221110-2111113333331133-1331301121030303"></a>

## Next pages — clear_secret_info / 212012113212 / 6

- [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101101112201311-0333221303333211-0112221100302133-3311121301221100-1010013033002001-3322312131132212-3002230003232113-1133102101132000"></a>

## palo_alto_fw_service.service_nodes — service_nodes / 303320330012 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- palo_alto_fw_service.service_nodes

<a id="canonical-3111023333233030-2330222103003112-0333322321203330-2233021223312021-1311003210112212-0201320111330121-3113223003301321-2332313012211021"></a>

Type: `"single"`. Computed.

Configuration parameter for service nodes.

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

<a id="canonical-1131033212003012-0112030102322311-0100111302113213-3113202332032010-3000022130333002-2033033220303200-1031231320222232-2120320203030313"></a>

## Direct properties — service_nodes / 303320330012 / 3

- [nodes](data-sources--nfv_service--reference--group-004.md#canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102): complete subsection reference.

<a id="canonical-2033322022032123-3220133212320331-1211313020221221-2230120220132221-1133300102301010-0032120333233212-0000010201033210-0331323213203122"></a>

## Next pages — service_nodes / 303320330012 / 4

- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131121222323222-3101303013212203-0012112130102222-3103311211203013-1133212300303333-2202010103003022-3310032003301111-3122232003103000"></a>

## palo_alto_fw_service.service_nodes.nodes — nodes / 321110133221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222)
- palo_alto_fw_service.service_nodes.nodes

<a id="canonical-2021220321123132-3121211221022002-2333302232122320-1322013320133220-0220222131220211-0123121131133031-2232100210132222-0123101033023021"></a>

Type: `"list"`. Computed.

Palo Alto Networks AZ Nodes. Configuration parameter for nodes

Upstream description:

Configuration parameter for nodes

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1310211300020012-2202333023202120-3130103030113231-3102200223101102-3302022121220313-2133120333021013-1231303203300321-1033110233102120"></a>

## Direct properties — nodes / 321110133221 / 3

<a id="canonical-1212203330302023-0131200333033302-0213212023322000-0132220302323231-0203230222323020-3311231002213323-0002311031211103-3033221232211313"></a>

<a id="canonical-0232211133302323-1303113300231211-0011310113130311-0201123333300201-1311313032112032-3333121020321030-0112011330031131-1220202013020230"></a>

## aws_az_name property — nodes / 321110133221 / 4

Type: `"string"`. Computed.

AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is
one of the AZ for sites.

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
    "pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  }
}
```

- [mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-0321303313302233-1320211102133320-0210022301030202-0110133100031201-0021112121011033-2112101021033213-2012103021120332-3302201223130130): complete subsection reference.

<a id="canonical-0113120213023203-2000323212100300-2010013111002031-3010330120110202-3103302100122102-0222121002102230-2002211002221013-3211202030012003"></a>

<a id="canonical-2012201003313001-2132201320200131-0303133021013322-1231303303002200-2200221322130323-0010313021223000-1002113023023131-1021123013200020"></a>

## node_name property — nodes / 321110133221 / 5

Type: `"string"`. Computed.

Node Name will be used to assign as hostname to the service.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [reserved_mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-0132100303021233-1212211320110310-2111310033033300-2311021023102210-0210221010100133-0113013011200300-3020130310231331-2220022010021303): complete subsection reference.

<a id="canonical-2332311312222320-0220120033302231-0003103323230323-3132122200022132-3202012110323011-2030210202102230-2320100011220113-1110233203200203"></a>

## Next pages — nodes / 321110133221 / 6

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-0321303313302233-1320211102133320-0210022301030202-0110133100031201-0021112121011033-2112101021033213-2012103021120332-3302201223130130)
- [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-0132100303021233-1212211320110310-2111310033033300-2311021023102210-0210221010100133-0113013011200300-3020130310231331-2220022010021303)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0321303313302233-1320211102133320-0210022301030202-0110133100031201-0021112121011033-2112101021033213-2012103021120332-3302201223130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111212231312222-1022003112030133-1313333113200130-2330210120100212-1121213223313013-3010330223122232-0310010032203102-0000312110001211"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet — mgmt_subnet / 122001011201 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet

<a id="canonical-2311321201203132-1310331112122323-2301020112322233-3333013021230213-2322031010121303-1331330201210223-2223210123211211-0123030122001331"></a>

Type: `"single"`. Computed.

Configuration parameter for mgmt subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-1012001021212021-2312001013203202-1001132123323301-2000012200321313-1212232103320203-1003231320031212-2010130112133130-1012130210002231"></a>

## Direct properties — mgmt_subnet / 122001011201 / 3

<a id="canonical-3232211322303120-0100222302302010-1030331011330122-2011131013200322-3313310112201001-1021103012230001-3012131113100320-3023012232113201"></a>

<a id="canonical-0031330320321122-0031332302332120-3132023210010233-1111030120120023-3120021000221022-1120102303022320-2020330030302122-0310303230321313"></a>

## existing_subnet_id property — mgmt_subnet / 122001011201 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--nfv_service--reference--group-004.md#canonical-0333002132210012-2200021101132032-3123221112201211-3312100131300130-0320231303320102-3302313232222033-3310301232233233-3110200230132112): complete subsection reference.

<a id="canonical-1311110233021200-1211020233303020-0101130230323220-3223011000330110-0230102223222321-2020123002101032-3323201310233232-2123133221213320"></a>

## Next pages — mgmt_subnet / 122001011201 / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--reference--group-004.md#canonical-0333002132210012-2200021101132032-3123221112201211-3312100131300130-0320231303320102-3302313232222033-3310301232233233-3110200230132112)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0333002132210012-2200021101132032-3123221112201211-3312100131300130-0320231303320102-3302313232222033-3310301232233233-3110200230132112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120020133323313-2312032323320022-3301001031011021-0223110111311100-3320311220323210-1102203001232103-2031013033301132-2320211322121331"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param — subnet_param / 321300001002 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102)
- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-0321303313302233-1320211102133320-0210022301030202-0110133100031201-0021112121011033-2112101021033213-2012103021120332-3302201223130130)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param

<a id="canonical-2001031322313022-2203210011032310-1001000230012200-2121011033212133-2202312233130121-1331310233120212-3323110113333130-1300102320322333"></a>

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

<a id="canonical-1111323201200122-3000010123032212-3102103333213131-3212000112303112-3111001111212002-3300100300203211-1230021133313010-3022213322003002"></a>

## Direct properties — subnet_param / 321300001002 / 3

<a id="canonical-1300000212301130-1002333033031311-1222113002221212-0211220010323011-3321313331333222-3201122133231203-2223201333103202-2000110311201121"></a>

<a id="canonical-3001322332300233-2122011321110123-2030233101130203-0011120210333102-2232100112312323-2201010321201301-3101001010111210-0133313033323133"></a>

## IPv4 property — subnet_param / 321300001002 / 4

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

<a id="canonical-2000101310023323-0121131101112203-2003111201103031-2130310320212031-2030300213231322-3311332131112202-2001233301313203-1032131021001311"></a>

## Next pages — subnet_param / 321300001002 / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-0321303313302233-1320211102133320-0210022301030202-0110133100031201-0021112121011033-2112101021033213-2012103021120332-3302201223130130)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0132100303021233-1212211320110310-2111310033033300-2311021023102210-0210221010100133-0113013011200300-3020130310231331-2220022010021303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331033010021101-1133301320302330-0302233212230221-2130302103232012-0321111130012011-1110012130332222-1121013230201110-2302220333111003"></a>

## palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet — reserved_mgmt_subnet / 131033301331 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102)
- palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

<a id="canonical-3222001131002220-2123300131030320-3233023211132200-1031230303203110-2102302012331332-1010303213302032-0113112311121212-3213000222020120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved mgmt subnet.

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

<a id="canonical-0320310320220202-3232323231310101-3131221331000033-1300110030330020-3202313301321231-2003021320213313-0330301122212313-2101331011120310"></a>

## Direct properties — reserved_mgmt_subnet / 131033301331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111333203103331-0001300220003022-0113132033102121-3010023102132023-3101231012333022-1021323332110011-1230013232302131-0111120120210301"></a>

## Next pages — reserved_mgmt_subnet / 131033301331 / 4

- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
