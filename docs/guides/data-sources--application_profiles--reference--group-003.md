---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-1312200232032122-1102033010233113-0210102323013113-2201101213033313-3031323020222112-0300121231031123-2131000313000003-3330130311132203"></a>

## uid property — udp_server_profile / 322112223122 / 8

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

<a id="canonical-2003212112102330-1200201200111132-3323323331211111-0310202322323233-1002111023221111-1130330201212232-1323231123032021-1013111013113111"></a>

## Next pages — udp_server_profile / 322112223122 / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231232222211323-1020132203233330-1203313120210100-1312323111003002-1021003322010102-1013023330033211-0001211011003311-3213120111031120"></a>

## virtual_server.https — https / 022003232311 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.https

<a id="canonical-3110112100311110-2102131302131202-3130311013113200-3302111333312320-0231310322001110-1202132200130320-3130103300220103-2113202112113123"></a>

Type: `"single"`. Computed.

HTTP profiles.

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

<a id="canonical-2233010113213000-3120232012230101-1232212121320112-0131012000313000-0322131313210212-0203220033210230-3003322301030223-1222233310323212"></a>

## Direct properties — https / 022003232311 / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3001003302112110-3303032312322223-2133130310311231-0131323012021332-0210112201333032-3112021303011312-1000213203131133-0232013202102203): complete subsection reference.

- [http2_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-0232211301312331-3031322222310010-2213321213322101-0030303001230311-2333130203133130-0231333112101010-0230220121312123-0300210202101312): complete subsection reference.

- [http2_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-0133103111122212-0030132222212100-1322222321332221-1120301021330131-2303022021222200-2011233020213313-1331022021303013-2011010231130030): complete subsection reference.

- [http_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1033020221112231-2312002311211110-1202132231000101-1221103303220301-0002022023110333-1300122031222333-3211212113103022-2332232210331000): complete subsection reference.

- [http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-1122232102120303-3221012223022302-3211320120010213-0130131002322323-0312123133102233-2312312301131020-2020103320100202-1133220133210101): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-3322023320211211-1331022213302121-0133203222332103-1030110033331121-2132310011310012-3201031220331313-1103000123001311-1303223101200210): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1232131212003202-3233200300332302-3020033131100312-0013102030313023-3233122321231022-1311133030120233-3112020100110211-2033312301303113): complete subsection reference.

- [stream_profile](data-sources--application_profiles--reference--group-003.md#canonical-0210002202230203-3332310302110202-1331311203103002-0130212013211202-2323203303013132-3322010323130333-0303211301321220-0302221122220103): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2202313221112330-1012021332020313-2032233030332021-3020110300212031-2303012031001112-1120200231113030-3001212033333233-3231212323223302): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2323210303130233-0021121130131301-1201021331112013-3231311301121012-2000003101021203-1210303003100330-0230010033103000-2210101312311203): complete subsection reference.

- [websocket_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2103323231301330-0000131213010232-3232000232201002-1333132013310011-1100032312012331-0013323330310330-1313023021002211-2002230001300010): complete subsection reference.

- [websocket_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-3012332000033101-1232030021020103-1233321320121310-1012010101221322-0111221200010030-0303213133221032-3302100321231030-3030101313212331): complete subsection reference.

<a id="canonical-3223203221130221-0112001330202221-1133333201333103-2221311131320201-0212132100303121-1212012030331002-1022021030011221-1222010013010233"></a>

## Next pages — https / 022003232311 / 4

- [virtual_server.https.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3001003302112110-3303032312322223-2133130310311231-0131323012021332-0210112201333032-3112021303011312-1000213203131133-0232013202102203)
- [virtual_server.https.http2_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-0232211301312331-3031322222310010-2213321213322101-0030303001230311-2333130203133130-0231333112101010-0230220121312123-0300210202101312)
- [virtual_server.https.http2_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-0133103111122212-0030132222212100-1322222321332221-1120301021330131-2303022021222200-2011233020213313-1331022021303013-2011010231130030)
- [virtual_server.https.http_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1033020221112231-2312002311211110-1202132231000101-1221103303220301-0002022023110333-1300122031222333-3211212113103022-2332232210331000)
- [virtual_server.https.http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-1122232102120303-3221012223022302-3211320120010213-0130131002322323-0312123133102233-2312312301131020-2020103320100202-1133220133210101)
- [virtual_server.https.ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-3322023320211211-1331022213302121-0133203222332103-1030110033331121-2132310011310012-3201031220331313-1103000123001311-1303223101200210)
- [virtual_server.https.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1232131212003202-3233200300332302-3020033131100312-0013102030313023-3233122321231022-1311133030120233-3112020100110211-2033312301303113)
- [virtual_server.https.stream_profile](data-sources--application_profiles--reference--group-003.md#canonical-0210002202230203-3332310302110202-1331311203103002-0130212013211202-2323203303013132-3322010323130333-0303211301321220-0302221122220103)
- [virtual_server.https.tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2202313221112330-1012021332020313-2032233030332021-3020110300212031-2303012031001112-1120200231113030-3001212033333233-3231212323223302)
- [virtual_server.https.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2323210303130233-0021121130131301-1201021331112013-3231311301121012-2000003101021203-1210303003100330-0230010033103000-2210101312311203)
- [virtual_server.https.websocket_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2103323231301330-0000131213010232-3232000232201002-1333132013310011-1100032312012331-0013323330310330-1313023021002211-2002230001300010)
- [virtual_server.https.websocket_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-3012332000033101-1232030021020103-1233321320121310-1012010101221322-0111221200010030-0303213133221032-3302100321231030-3030101313212331)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-3001003302112110-3303032312322223-2133130310311231-0131323012021332-0210112201333032-3112021303011312-1000213203131133-0232013202102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233010023111012-2300023010002302-3210230131231310-0310232320022010-2031103211023231-0133213122001313-2000033332013102-3232120023032232"></a>

## virtual_server.https.client_ssl_profile — client_ssl_profile / 331012202203 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.client_ssl_profile

<a id="canonical-1132322120022333-0312030202333221-3321011032223110-3301111312201221-1130120122132222-1002223013031310-0302311312132232-2230003122102313"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1203232202303210-3102130320221013-2310101302221230-3212133031120012-1230021103001011-1002310301112133-1233220023331103-0000232322230122"></a>

## Direct properties — client_ssl_profile / 331012202203 / 3

<a id="canonical-1110121312130230-2001103131200203-1002102313133131-3303021103132200-0130032231211223-1300220112302302-1310022311212221-1003213311001200"></a>

<a id="canonical-2303201121313303-3020032100032230-3232212200303303-2331221201111300-3020031303010021-1120200132200322-0120303031113303-3123201211112300"></a>

## kind property — client_ssl_profile / 331012202203 / 4

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

<a id="canonical-3213033313130003-3221300102030120-0100313221202330-2311332320111010-1102210020102132-3000121221120320-2031021232010320-3210222111211002"></a>

<a id="canonical-2213320220312303-3210012032011012-1203230331131201-2001131101122223-1233222101220303-0301223021112302-2301232120211010-2022202320303303"></a>

## name property — client_ssl_profile / 331012202203 / 5

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

<a id="canonical-2033311002330301-2231010213033013-0030022112033122-3332110123102221-0113120203033013-1322003033232020-0311011023132232-3103020330101200"></a>

<a id="canonical-2213300211023111-2110222312010303-2000230033232230-2002300323211132-0110020321130301-3131122110110322-1000123202322233-3012102111130012"></a>

## namespace property — client_ssl_profile / 331012202203 / 6

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
  }
}
```

<a id="canonical-0112111010121310-3102223002232002-1203013211000031-0230311212231010-3021331010111232-0011332030130232-0100211101100303-1331233201010010"></a>

<a id="canonical-3301312103301022-1101001223023122-0220222023310010-1030332311230200-1312000200011012-3310020320121300-3121102023101010-1202320020302223"></a>

## tenant property — client_ssl_profile / 331012202203 / 7

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

<a id="canonical-3102103202302101-1123202102012103-3301002133331202-2120011031320023-1331100003303233-3322102211013001-3233002212032231-3020020120022222"></a>

<a id="canonical-2211221202311022-2103300120001120-0030221123221212-0030223321333013-0300010311002303-3201103303010322-1311031312323113-2011230132220023"></a>

## uid property — client_ssl_profile / 331012202203 / 8

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

<a id="canonical-1203221112203000-2210333212223002-1023213212123010-1213330002211323-1330110320302102-3333130312121201-0303112013001110-3133112021033333"></a>

## Next pages — client_ssl_profile / 331012202203 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0232211301312331-3031322222310010-2213321213322101-0030303001230311-2333130203133130-0231333112101010-0230220121312123-0300210202101312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021313332003320-1213232023123211-1213000032200011-0102203321021223-1122320333213211-0230023133122032-0003001322230100-0133122230330113"></a>

## virtual_server.https.http2_client_profile — http2_client_profile / 333323331230 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http2_client_profile

<a id="canonical-2300311230302121-0332301030310130-1130120012100320-2113332311032112-0221030233102130-1111102222010113-0033022301021033-2213113030330110"></a>

Type: `"list"`. Computed.

HTTP/2 Profile Client. Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3110121111221303-2021222001013221-3201300110323123-0203022233031322-2202001101301213-3103322021013333-0322010323210301-2110231130312300"></a>

## Direct properties — http2_client_profile / 333323331230 / 3

<a id="canonical-2203030312232013-3121232202213222-0130200302212031-1332233000303222-3133213310313232-1320020223300131-3103231333212211-2131010323301103"></a>

<a id="canonical-1222233013210132-0121132313220013-0113202123330322-1132113030332230-0111113033231322-1302132302101333-0222101221033333-3322312220003212"></a>

## kind property — http2_client_profile / 333323331230 / 4

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

<a id="canonical-3020012020102033-3031123323232103-1101332333121031-3310203103033000-0323113313003102-2231033220100231-1120320202220231-0021100003102003"></a>

<a id="canonical-2233330333001123-0010220023323322-1202221012310330-0300323212103222-3133110013033333-0312023201021321-0011112232302130-1322231200131232"></a>

## name property — http2_client_profile / 333323331230 / 5

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

<a id="canonical-2031100001000231-1032101003301222-1130121203102332-2132333023131010-1313101131303212-3333120331003113-1210022321033001-2310122203121332"></a>

<a id="canonical-0120233102112323-2333030210030003-1122232210330120-3223231322111222-1313103302322020-0323210211302123-3110110130132233-3203210113303130"></a>

## namespace property — http2_client_profile / 333323331230 / 6

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
  }
}
```

<a id="canonical-2102323113323032-3303022213303210-3110011201110201-3102333032000120-3200233312023010-0222202311102321-3321200121212102-2331311230020032"></a>

<a id="canonical-3021300030000212-3012000110112131-0033023221312211-2103002220331331-3013130231101130-0210003120233301-3010300322202213-0312202232001010"></a>

## tenant property — http2_client_profile / 333323331230 / 7

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

<a id="canonical-0113223121300303-0230023313301202-0133133133313111-2113020333313130-3011122232313030-2032023212113113-3100203330213232-0200001101003223"></a>

<a id="canonical-0202200312001012-3032121113210022-3222302133002301-0233200320300100-1332032310031123-3032121230030320-3321013123303130-1102223200201010"></a>

## uid property — http2_client_profile / 333323331230 / 8

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

<a id="canonical-2333213130330100-1013100232232102-0113111030121230-0212300223110232-3221132013222231-1211023320323023-1323203222102112-0133201322232011"></a>

## Next pages — http2_client_profile / 333323331230 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0133103111122212-0030132222212100-1322222321332221-1120301021330131-2303022021222200-2011233020213313-1331022021303013-2011010231130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130302031110013-2131022202013123-1220113300013122-0230033312233100-0320102101301222-3102201123231303-2210212213300010-2322021202312303"></a>

## virtual_server.https.http2_server_profile — http2_server_profile / 130310203010 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http2_server_profile

<a id="canonical-3032313001230132-3323121103023232-0030300210011030-1130303023121331-2203133002010120-0303200302222011-0302021130001313-3122222101303101"></a>

Type: `"list"`. Computed.

Configuration parameter for http2 server profile.

Upstream description:

Configuration parameter for http2 server profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2000100200110220-0232123021312111-1111030002201120-3230332323110231-0301113123103231-1031102112002120-3203003103130203-3220013301321200"></a>

## Direct properties — http2_server_profile / 130310203010 / 3

<a id="canonical-1202301132213331-1203021030233232-3130223333223111-3232102130330231-3032012001222110-2030210230301101-1100131103020110-3201103013230121"></a>

<a id="canonical-0003111100021321-2221030332121011-3022211212230113-2231202110332210-3313123020020301-3330102300012121-3131320100000021-0021021130202133"></a>

## kind property — http2_server_profile / 130310203010 / 4

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

<a id="canonical-3223330231301021-2100323011202103-0303120122312233-1130131121332113-3302210111222212-0310231232121113-0001022121001002-0313100230201000"></a>

<a id="canonical-1222013203300002-2221131212112303-2302302211111331-1002201213321110-2130232030222120-0011322212213113-0231203230300003-2232220221132232"></a>

## name property — http2_server_profile / 130310203010 / 5

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

<a id="canonical-0033232312013133-0001112331212301-2033100130301223-3023303120230001-1331320012111232-0222311310203001-1103012201231303-3102002200100102"></a>

<a id="canonical-3313230013311232-0130100301330031-3110132001320113-1023031122013122-2310013031130002-0332023313321203-3230201101032132-1203113022100320"></a>

## namespace property — http2_server_profile / 130310203010 / 6

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
  }
}
```

<a id="canonical-3121201231013100-2201323213221001-2113310322022100-0232323033123122-1232133331201202-2133003212303330-2221320121331011-0100133130003133"></a>

<a id="canonical-3012101130201133-1130301210123103-0133010303112122-2022233033012231-2031011000322221-2312310223301312-3010302300302311-3023011211310320"></a>

## tenant property — http2_server_profile / 130310203010 / 7

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

<a id="canonical-0002203310232131-0133132223210032-3321200332310202-3002203210020131-0203032211210212-0000210210003031-3011322201110133-0231231033200011"></a>

<a id="canonical-2111133113121111-0221333003030213-3023223223111320-3330223032022202-0303320013010323-1200133130132003-2002210132202120-1031303303120231"></a>

## uid property — http2_server_profile / 130310203010 / 8

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

<a id="canonical-2233212012022222-2220123122213231-1222132211311022-0103202211133111-2112313310133310-0310330231002301-0311021200111310-3011100113131133"></a>

## Next pages — http2_server_profile / 130310203010 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1033020221112231-2312002311211110-1202132231000101-1221103303220301-0002022023110333-1300122031222333-3211212113103022-2332232210331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012102133232013-1333323321310000-0223230021200121-3021332210332221-2322331100320323-0102221133110100-0121311013323011-0120030212030212"></a>

## virtual_server.https.http_client_profile — http_client_profile / 321102002012 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http_client_profile

<a id="canonical-3131300113300020-0011113001200100-2121210322120013-3211213213122233-0003210302021133-0330020220231003-3331332200203012-1213322003103003"></a>

Type: `"list"`. Computed.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3222213122031220-1231232322321202-2102000222322332-0102131211312111-0220212231003032-3231330130233310-3321312313333033-0212333203213022"></a>

## Direct properties — http_client_profile / 321102002012 / 3

<a id="canonical-0311231031223312-1003233110001021-1330321330102201-3310001122320102-0233220002200231-1021012000031103-0223003332101120-1133323203202031"></a>

<a id="canonical-2231010223313220-1321102200003110-3033313300111230-2230202221122010-0220023120210031-0132310031122323-3213131130021123-1232010133223301"></a>

## kind property — http_client_profile / 321102002012 / 4

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

<a id="canonical-3210120202130120-3122200123122223-2101031210003201-1231020113131103-1120112130330211-2313010312302330-0223121021321122-1120033321301012"></a>

<a id="canonical-3321211120132301-1133100233032000-3000220021221321-1030101123313200-2211202102202032-1300323112033132-2323011131030312-0302213331310313"></a>

## name property — http_client_profile / 321102002012 / 5

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

<a id="canonical-0113320320112000-2022020213132233-3201323212030132-3212202201011120-1021323133120231-1130330123031211-3123230000032131-1133011021322211"></a>

<a id="canonical-2000233013121311-0011201302023113-0020013222002333-3301012312023332-3230232122130221-0132130013111111-1212220202233021-0203211331103203"></a>

## namespace property — http_client_profile / 321102002012 / 6

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
  }
}
```

<a id="canonical-3013122132032220-3300200112313121-2133301131222123-3100233121013313-2331231100323012-3101110000001020-3023333113213212-2101030100013002"></a>

<a id="canonical-3322011332202210-1301220310121331-3310223232023202-0312213301223210-2123102232003320-0123133122300022-3321121212110201-0300223330122211"></a>

## tenant property — http_client_profile / 321102002012 / 7

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

<a id="canonical-0212230031010012-0201111222230021-1302231113311132-3133123131023130-0021121001231123-1132301212233032-2112022201202112-0000022132301230"></a>

<a id="canonical-1210301001132233-0303123021103211-0212223301010111-1110331030131212-1213212032032220-0031133111332103-0110020120303001-0000103021110311"></a>

## uid property — http_client_profile / 321102002012 / 8

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

<a id="canonical-3121333002222312-3303120311320122-0131322000130312-3203321011321302-2232012233310032-1011211201122110-2110312311111113-3110221023022301"></a>

## Next pages — http_client_profile / 321102002012 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1122232102120303-3221012223022302-3211320120010213-0130131002322323-0312123133102233-2312312301131020-2020103320100202-1133220133210101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333200222021020-3000330033323100-1002213113320313-1033110302312100-2223031030331110-2211111022023031-2233231101313022-0320231123102000"></a>

## virtual_server.https.http_server_profile — http_server_profile / 020200121010 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.http_server_profile

<a id="canonical-3132302132312210-2020323020222112-0001200000121223-0212232330220030-0003202223331003-2221331033000110-3033303120211121-0231001312112020"></a>

Type: `"list"`. Computed.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3030202103020113-3323013313231301-2030321122221201-3210000301231010-2113321022312310-0003130010200011-0222131231131330-1003202322131010"></a>

## Direct properties — http_server_profile / 020200121010 / 3

<a id="canonical-2313123202213030-3331122332020222-0231131311022311-2030221320303210-3111033333222111-2301100010102000-1303122020202000-1201232312322311"></a>

<a id="canonical-3330002302102001-0032122020213002-2111102301011122-3123033021012121-3321302313303031-3003122233010231-0202010333022003-0313013111023232"></a>

## kind property — http_server_profile / 020200121010 / 4

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

<a id="canonical-1033133132012220-0313302002231332-0322332232330211-3013130103131212-3133331010033001-2113322113310302-3321011000102031-2003003101331213"></a>

<a id="canonical-0220221022233001-0000331233331310-1201020113312221-3213213013102310-3101023131230200-0220303222023002-2013030031110133-1332302102232212"></a>

## name property — http_server_profile / 020200121010 / 5

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

<a id="canonical-0203121001132023-0203020112110320-3102230001231011-2233302203012112-1031230303112010-3202203103000223-0213013312323102-0101212220333122"></a>

<a id="canonical-2320132310120013-0023222011310032-1333021222010130-2033031221223020-2333303300010322-0031122033132201-0100211210222313-3023100103012000"></a>

## namespace property — http_server_profile / 020200121010 / 6

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
  }
}
```

<a id="canonical-2210223023312222-3232112103301023-3000033211200110-3223131332100020-3001303110110331-0003102012120113-0102333211110310-0112210233032223"></a>

<a id="canonical-1103323233222013-0221331210113011-0001111213031130-0323222320103232-2221210122221011-0301001013131301-3201133033012202-0023121232231311"></a>

## tenant property — http_server_profile / 020200121010 / 7

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

<a id="canonical-1012232303223130-1311121212303020-0211321020020313-3000230031033011-0123023111100110-0102021101333001-0011130133022030-1212310110012030"></a>

<a id="canonical-2322223211031122-2301332103000230-3032023123232301-3010023113312111-0221221013121333-1000221100011013-2222022023330012-1122201311113023"></a>

## uid property — http_server_profile / 020200121010 / 8

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

<a id="canonical-2100103002132003-3200020101200121-2201103133022101-2223320033300131-0103022121313032-3321213212310223-1101031000231001-3330021202330223"></a>

## Next pages — http_server_profile / 020200121010 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-3322023320211211-1331022213302121-0133203222332103-1030110033331121-2132310011310012-3201031220331313-1103000123001311-1303223101200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332010133201331-1312030103311220-1001010010112113-1323313010200010-1021310310111310-2333222101113212-3320110122211212-0110133313223320"></a>

## virtual_server.https.ocsp_profile — ocsp_profile / 100102031332 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.ocsp_profile

<a id="canonical-2013002101022013-2331200321301103-1321203302300002-0113011110003003-2233212212030320-2102000221123111-0330031003132102-1112012323101021"></a>

Type: `"list"`. Computed.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2210011303303033-1100220021012233-3032321001221002-1331211000323211-0002213131102200-2321332110321020-3130133321011313-2113031112231011"></a>

## Direct properties — ocsp_profile / 100102031332 / 3

<a id="canonical-0333132003010323-1120331001331123-3310322213323130-0001032331212300-1133120131302323-3033121313311120-0211223222112230-3230022231333202"></a>

<a id="canonical-1222210310212132-3130212232222013-1010232031222002-1020121211322212-1301210202322131-3203320222010001-3033330212210302-1321231221303022"></a>

## kind property — ocsp_profile / 100102031332 / 4

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

<a id="canonical-2230023313120100-1330310233232010-1320231330332233-0103000301232031-0130023113011302-2232132112201320-1012220320013001-1032232031122212"></a>

<a id="canonical-0223132011022022-3211023320222222-1122202330232110-2002002123132010-3031020021003222-0102131001213202-0121221030310110-1201310200102211"></a>

## name property — ocsp_profile / 100102031332 / 5

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

<a id="canonical-0331122031032201-2023303021010123-2333311332200113-2011112031123032-3302313123333113-0012233132033133-1030012303030221-3220113120200330"></a>

<a id="canonical-0130021313001010-2022023221230110-2111002312000003-3231113001223012-0022031332110313-3013323311210313-1311102113123003-3100021300100132"></a>

## namespace property — ocsp_profile / 100102031332 / 6

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
  }
}
```

<a id="canonical-3322233310021320-1312220233121113-1020031232221023-3201103031032002-3210033330122110-1013321200010323-0000100322132202-0013012333221102"></a>

<a id="canonical-3012133101322031-1332112121002011-3023030233101113-0201012030230003-1302211331300022-1212002310012323-3132032321111113-2021033112303133"></a>

## tenant property — ocsp_profile / 100102031332 / 7

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

<a id="canonical-3110303230120211-2032002231223301-1320102101230200-2002000300101110-1110113211211010-0030300222021022-1222112133013200-1300000130031200"></a>

<a id="canonical-3121001333213033-1012313201110312-3300310010221133-0231200323123120-2221320231011023-3012000131203120-1222202111130212-0130130221020130"></a>

## uid property — ocsp_profile / 100102031332 / 8

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

<a id="canonical-2312303212113201-1132213002220303-3300302333112213-2230221310310322-0312222212210201-3332211302123000-2233202331222222-1233201313300023"></a>

## Next pages — ocsp_profile / 100102031332 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1232131212003202-3233200300332302-3020033131100312-0013102030313023-3233122321231022-1311133030120233-3112020100110211-2033312301303113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320020203212102-0101312111203103-0022001131312201-0331322332330322-0132211201311331-3130101313310003-0221320023022023-2001033332220000"></a>

## virtual_server.https.server_ssl_profile — server_ssl_profile / 102310001203 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.server_ssl_profile

<a id="canonical-3303311223330312-3033332311203113-2213312330332231-2300030212111132-2232201300002011-2011301203332012-3231333312121132-0112202100122202"></a>

Type: `"list"`. Computed.

Configuration parameter for server SSL profile.

Upstream description:

Configuration parameter for server SSL profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1330211202233311-3123011131120313-1120101131310110-0211003233322222-2230033021300021-2232130012330301-2332021033031132-2212203112301021"></a>

## Direct properties — server_ssl_profile / 102310001203 / 3

<a id="canonical-2003120003221002-1223130000200002-1311221231030330-3323013323220100-2001121132220131-0102313123312332-1001220333133001-3021133133120310"></a>

<a id="canonical-2311100112230213-0011132232302232-2330012233111102-2203221102223332-2200020120013320-3301300111223123-1001101202313311-1023103323331203"></a>

## kind property — server_ssl_profile / 102310001203 / 4

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

<a id="canonical-1323230312222023-1102033003233303-3033113201322003-2301312203332210-1220230321310322-2031333012333011-0321233332123113-3101220303303233"></a>

<a id="canonical-2102220132331322-1032310113033111-0232102310101130-0220000320233011-2313020212302123-0300302213302013-3112110013211100-2131020331203212"></a>

## name property — server_ssl_profile / 102310001203 / 5

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

<a id="canonical-0000120220131302-3312322122131013-2230331032101122-2001222121232320-3223201220221111-3231121103122201-3332022131220100-2300223323101021"></a>

<a id="canonical-0230321322200220-0031311120331123-3313203221033122-0120103102100103-3212223020131121-2322122100213030-1120322313111302-1200131303002120"></a>

## namespace property — server_ssl_profile / 102310001203 / 6

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
  }
}
```

<a id="canonical-3203330132232233-1123323113131011-2211210121320212-0102303120111012-3200212011203330-3122333003123130-1223213321023120-3222101303211133"></a>

<a id="canonical-2101103101322232-0320033331032023-1202202110110321-0210013322122211-3000012201131312-2322021013331022-2001213032033000-2023130122102320"></a>

## tenant property — server_ssl_profile / 102310001203 / 7

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

<a id="canonical-1131310220201110-0200002213220302-0211122122322001-3121111100222010-1021010002120130-2200013233333033-2011122113132300-2300020121320232"></a>

<a id="canonical-0232131212020312-0012203010113320-1201021111111302-3020301100233310-1011211131200213-3331311021230323-3113311001131232-3023013201101331"></a>

## uid property — server_ssl_profile / 102310001203 / 8

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

<a id="canonical-2200222301320100-3321021221000331-3222221023201021-0312012202003012-0200113311200221-2301130202033313-2020220303332200-3023031233021120"></a>

## Next pages — server_ssl_profile / 102310001203 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0210002202230203-3332310302110202-1331311203103002-0130212013211202-2323203303013132-3322010323130333-0303211301321220-0302221122220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030221133033023-3022123221220021-1323120313213023-0120000121313230-0130220112020301-3033230130302012-2231203013011132-0212333313131123"></a>

## virtual_server.https.stream_profile — stream_profile / 021132203122 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.stream_profile

<a id="canonical-2101112123203110-0231221321022130-3103201012031200-2300212020012011-1230032322101301-3311322301203000-1102100102000233-0331221321313011"></a>

Type: `"list"`. Computed.

Configuration parameter for stream profile.

Upstream description:

Configuration parameter for stream profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1332201321302103-1300110200313332-0032011301000102-2032333321302130-3313323110313223-3300301032030210-0020032221301201-0313222112310301"></a>

## Direct properties — stream_profile / 021132203122 / 3

<a id="canonical-2121222122231220-3003121210123321-3022213100303002-2112102222210331-0122110232203201-2020110010230013-3310202321301001-1100211213332302"></a>

<a id="canonical-1221333120210230-1232210130321330-3211323113330221-1322310002130000-0200010202210221-0233300210001110-0010320022321111-2203323201011213"></a>

## kind property — stream_profile / 021132203122 / 4

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

<a id="canonical-1221211022022220-2320212233303101-1212300121102300-1321012331102123-0003011221003303-1033023031111032-3123122100121202-2200202113010331"></a>

<a id="canonical-2011013332011221-0102030200102113-0311313033313032-3012333213330023-2112330230220100-0030122000202010-2221302022301031-3310011232323313"></a>

## name property — stream_profile / 021132203122 / 5

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

<a id="canonical-1323301301310000-2200120032213330-1233200232021002-3102101333111003-0032113310102103-3322013311201003-1333210302001003-3210030121102231"></a>

<a id="canonical-3221310333002121-1121202101220113-1030122202131133-3223122103012011-0303303011213213-2012110013322211-2321300123123103-0330313303310101"></a>

## namespace property — stream_profile / 021132203122 / 6

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
  }
}
```

<a id="canonical-3112011222113010-1100200331000332-0300133333112322-3020010302200310-0022133213200113-2022023232322223-3023223231113201-0211020312011022"></a>

<a id="canonical-3003313021302110-0331113112121121-0023113210012132-3122211002000313-1133031102113103-1303210010103121-3230032032003011-2222022310332100"></a>

## tenant property — stream_profile / 021132203122 / 7

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

<a id="canonical-0032123333333230-0013230113220010-1313011122213332-2220233113113320-2323333000130303-3202120010003011-2211030223311031-3102301001030012"></a>

<a id="canonical-0020202230202330-3002020330102313-1123012221210101-3220230011311300-3300011220322021-1302223130220111-3203012132220121-1300110320012231"></a>

## uid property — stream_profile / 021132203122 / 8

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

<a id="canonical-3012333013030312-3332201233203112-3312311201123120-1122230313113233-2323123210233103-3211311231210022-0121232302223111-1031202002123202"></a>

## Next pages — stream_profile / 021132203122 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2202313221112330-1012021332020313-2032233030332021-3020110300212031-2303012031001112-1120200231113030-3001212033333233-3231212323223302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333233132321111-2332011013120021-1100010102330212-3322031131221132-3032210221123331-0103130001130311-3311023323332303-0202212020230021"></a>

## virtual_server.https.tcp_client_profile — tcp_client_profile / 020231221321 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.tcp_client_profile

<a id="canonical-3311100221020133-0023232010132211-2010011331121213-0031000022333011-2010010233030320-1311003313310322-2210321223120132-2003213033101022"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1123201212233323-2322032110211301-2220302033110202-0230120233333011-3111320233131013-0220311212321210-0301030000120113-2000032213210130"></a>

## Direct properties — tcp_client_profile / 020231221321 / 3

<a id="canonical-3000013103023010-2330212310323033-3331201020222200-2300032002011121-1030111032321030-3123233320330320-1120003212313332-3223330333223231"></a>

<a id="canonical-0123201033311200-3121230032111102-3021022220311100-3221011031102210-1202020123001223-2332331303033021-1322212113113321-1000110122102132"></a>

## kind property — tcp_client_profile / 020231221321 / 4

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

<a id="canonical-1103011220011233-0321202113110021-1233000302213123-1330032120031301-1032103112102200-3213311102120310-2131001220211321-1312231201322121"></a>

<a id="canonical-2021103222202101-1112120123210223-2330233132000222-2012113131032300-3212322300030112-3222213303130122-3311001231333012-2231321313123100"></a>

## name property — tcp_client_profile / 020231221321 / 5

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

<a id="canonical-3232010022121310-1211101223223122-3211031002300112-2221020122011223-2223232313022020-1000003023320233-3332211113001112-1133001100310232"></a>

<a id="canonical-1001123113122001-1031300321320201-1201211302332023-3131032003231301-0210013022330021-1322333133312032-0200230223000021-1322321123312330"></a>

## namespace property — tcp_client_profile / 020231221321 / 6

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
  }
}
```

<a id="canonical-1010202122201212-3223101002323311-3323013011303332-2201003313100330-0002331213203202-0210133211322131-2013211312212333-0010231013133300"></a>

<a id="canonical-3300222002131323-2023212132100301-2331002022100230-0231233200231203-1303213101203103-0302121001302332-2032230301003133-3330310233311211"></a>

## tenant property — tcp_client_profile / 020231221321 / 7

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

<a id="canonical-0020322001231033-0020201001103000-1130013231123021-1213213011110002-2313300100001021-3202223011030302-1003323133110233-2100012113311032"></a>

<a id="canonical-0101201112220232-1302030220012302-0223223303033331-3332103131112212-1012211300223320-1100003310103332-3222202302303012-3122223133320332"></a>

## uid property — tcp_client_profile / 020231221321 / 8

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

<a id="canonical-3312133333213201-3023022311311113-1012121333100030-0323000201320012-2331300310210201-2311031223322120-2333333333001032-2331211203231211"></a>

## Next pages — tcp_client_profile / 020231221321 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2323210303130233-0021121130131301-1201021331112013-3231311301121012-2000003101021203-1210303003100330-0230010033103000-2210101312311203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220202023003112-0030212110020111-0202012132021112-1312222210222201-3021330333121202-1333021012131310-3012012222332313-0102003121103130"></a>

## virtual_server.https.tcp_server_profile — tcp_server_profile / 323002003301 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.tcp_server_profile

<a id="canonical-2333122203230102-3331000211011113-0031023332320321-1200000031100101-3333133213321022-0131033020022213-2301223321102130-1311030113100202"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0121033123210223-2130010013032031-3121010301030201-2101311301322300-0333220323033002-3230301112030132-3033332111312021-1313212003013302"></a>

## Direct properties — tcp_server_profile / 323002003301 / 3

<a id="canonical-0013012323002312-2102331123231221-3321222103101202-3223212101113302-0001220321120300-1313303203230220-3100212321213302-2032130012332223"></a>

<a id="canonical-0302113233122103-1333231020113313-3101211120231011-0113101103330132-0202133331103212-2113100001121222-3221300210102112-1020220210033110"></a>

## kind property — tcp_server_profile / 323002003301 / 4

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

<a id="canonical-0101010133201312-1311303112221210-0313101110330203-0121102123023200-0311013002111200-3200231221011232-3330311203333100-1000320201312101"></a>

<a id="canonical-1102001223003030-1222213301300012-3300223123313031-0232213013002003-1130321301230132-3231211302001301-3102310223113330-3200331131010022"></a>

## name property — tcp_server_profile / 323002003301 / 5

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

<a id="canonical-3213110330021022-1300303210103311-1200130123220003-2223303111133100-0022030323301221-3303203201230303-1100221103300031-3132030201310202"></a>

<a id="canonical-1111022020112023-1033303323330032-3301102301030200-0120011230022001-3212112330201131-0123130330132110-2310333131012123-3122112000222130"></a>

## namespace property — tcp_server_profile / 323002003301 / 6

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
  }
}
```

<a id="canonical-3332321010023303-2101121321323103-3022220032202230-3223230130111301-2311310131032001-3123030223220033-3010200221002202-0203201330120022"></a>

<a id="canonical-2312131113220200-0131011020111002-1130032303333032-2233210321333310-0022321233211132-1103312313321213-0320111303312111-2201330322331320"></a>

## tenant property — tcp_server_profile / 323002003301 / 7

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

<a id="canonical-1223313010313313-2333222130310231-3023011022002022-3203010033002221-0201112013000013-1102301122020330-3111031123131230-0200330302303321"></a>

<a id="canonical-2011111102033012-0332332020210232-0123322300021023-1313333320112132-2031121222120223-3030123131211123-0022011331331231-3010313133011012"></a>

## uid property — tcp_server_profile / 323002003301 / 8

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

<a id="canonical-2202002131222320-1001222112130132-2312332023331122-0021223031111133-3101321210212221-0323212032200230-0210322213013132-3230310300220003"></a>

## Next pages — tcp_server_profile / 323002003301 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2103323231301330-0000131213010232-3232000232201002-1333132013310011-1100032312012331-0013323330310330-1313023021002211-2002230001300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012233031121320-0232311032320201-1132331211023020-0023003120023122-1012132133133002-1313033120300301-3232200233031331-0013202201000323"></a>

## virtual_server.https.websocket_client_profile — websocket_client_profile / 220321213321 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.websocket_client_profile

<a id="canonical-3110112022222213-3031103003231113-2210033100330223-0310132300103132-0223331232220230-0023202223300301-2213232333332010-2300321331131331"></a>

Type: `"list"`. Computed.

WebSocket Profile Client. Web-related configuration

Upstream description:

Web-related configuration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1112033120223212-1113320133101111-2212220320121321-0023312231133220-3210000211111120-3321332221021222-3210133332000310-0322120011131233"></a>

## Direct properties — websocket_client_profile / 220321213321 / 3

<a id="canonical-2131032331210311-0021210032233031-2212110110323130-2003022003203300-0220211333330021-1213012133012212-2313323222131102-1200120200323202"></a>

<a id="canonical-2103031200020312-0312213312212332-3230030330121131-3313301032011210-1022210030000003-0112313032321022-0333020230201302-2231010211122112"></a>

## kind property — websocket_client_profile / 220321213321 / 4

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

<a id="canonical-3330323230322020-2012031130030030-2113120330302100-0213200231332113-1013003120101320-2233132220023213-2213003311230130-3133031110103210"></a>

<a id="canonical-3200101033020033-2000122232030301-1331111211202123-2312012013213311-0103203330021223-1000222102013203-3302000333003221-2301102131210031"></a>

## name property — websocket_client_profile / 220321213321 / 5

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

<a id="canonical-0110203121212333-2032333011013122-2231111302201211-0023122303121011-2120131233032302-3100031303021212-2210132220213330-1231322001022102"></a>

<a id="canonical-1212110010332013-0131131210113222-3112220001121221-3332221123003300-2203030103311010-1020121323022130-0302121323112200-0213003112322203"></a>

## namespace property — websocket_client_profile / 220321213321 / 6

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
  }
}
```

<a id="canonical-1010103130330300-3233000010210011-2300111333231231-0002203032303110-0323123333311023-3210313102301312-1222113131321323-1003020200223120"></a>

<a id="canonical-0122101000002020-3120230233201030-1011120302131113-0010032021002313-2021220220223331-3320301132323321-0312031331233003-3120020231121332"></a>

## tenant property — websocket_client_profile / 220321213321 / 7

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

<a id="canonical-1303232103101303-3000223033320023-1032313332322333-2002313232312132-3023100201120110-0112023020122122-1032321302023332-3332313012112301"></a>

<a id="canonical-2220310302331322-1123030232030020-0233122311302332-0102013000210223-3101010111133332-0202220110000320-3110030003232030-3302000322001113"></a>

## uid property — websocket_client_profile / 220321213321 / 8

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

<a id="canonical-1232232302301110-1121300032122120-1213130323303213-0023320201312233-2102210221123122-1323022213310130-2113130130123203-1323002022220030"></a>

## Next pages — websocket_client_profile / 220321213321 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-3012332000033101-1232030021020103-1233321320121310-1012010101221322-0111221200010030-0303213133221032-3302100321231030-3030101313212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220200011322113-3330101210313322-2202223320022102-2301302310031031-3103122100211023-2210311122110100-3202102300111320-0210101133112100"></a>

## virtual_server.https.websocket_server_profile — websocket_server_profile / 020103201231 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- virtual_server.https.websocket_server_profile

<a id="canonical-2012111301231322-1111322030030322-2111202200121021-1030233113133301-1101212311113022-0101312003203001-0210220000030302-3210332010000033"></a>

Type: `"list"`. Computed.

WebSocket Profile Server. Web-related configuration

Upstream description:

Web-related configuration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2322103002211333-0110011132103030-3333001122033212-0200233222301202-3132213330233113-3312131333112113-0301310002122110-0321220012003331"></a>

## Direct properties — websocket_server_profile / 020103201231 / 3

<a id="canonical-0113033232120111-1002201022313210-1030311130010023-0330023210033321-0321320202020100-0113311232230321-0011102320232210-3231213010232232"></a>

<a id="canonical-2013323111320222-3010022332312100-3213211320312233-2202210131101221-1031312101310133-3031010230303313-1230120123003321-2221203213033213"></a>

## kind property — websocket_server_profile / 020103201231 / 4

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

<a id="canonical-1021000021011333-1233331232122330-0220113002301313-1002332003100013-0122220032220202-1032332031021002-1332013333133021-1232312100301023"></a>

<a id="canonical-0333123013130223-3033303331321320-2321011100103220-3323121312111302-0013133112323303-0013011032211300-3020233202302120-1312302222003322"></a>

## name property — websocket_server_profile / 020103201231 / 5

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

<a id="canonical-1231220022320220-3302223200010122-0200303000022131-1112222231322113-1203111323210202-0101212310121011-1110132012210213-2303211230123023"></a>

<a id="canonical-1022330313111331-2031100302003202-0022030323003003-1101023323113321-1303032300011001-2203200302330331-0333321320212023-3132111110130021"></a>

## namespace property — websocket_server_profile / 020103201231 / 6

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
  }
}
```

<a id="canonical-0012033122103103-3012233012033213-0001202231333010-0301331321100133-2002120120102011-2212122300310322-2123102231131002-1031133331211032"></a>

<a id="canonical-0113123030110323-3303330321201200-2100213122300211-1020102112002103-3221011113111131-1311001132231032-1031202310231203-0033111201301131"></a>

## tenant property — websocket_server_profile / 020103201231 / 7

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

<a id="canonical-0200303012122222-2213200013030120-2112010231302332-3023210330121221-1212123022222013-2200211110011231-0331213032113332-3031132301122001"></a>

<a id="canonical-2011022301033112-3332311310023333-1121222023303000-3301000103222013-1031202112033022-2111013132032323-0020221102132122-3020233031301213"></a>

## uid property — websocket_server_profile / 020103201231 / 8

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

<a id="canonical-0133313122201320-1122031103112300-3332112223213121-0212023313312112-0010231321312012-3132213330210301-0031212222231133-3110020201332212"></a>

## Next pages — websocket_server_profile / 020103201231 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111202103020311-3330310301311111-2320303001030231-0223030332233222-0322110122003112-3010002001331131-2010313223221212-2210133010323321"></a>

## virtual_server.immediate_action_on_service_down — immediate_action_on_service_down / 000132012120 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.immediate_action_on_service_down

<a id="canonical-0210331121211012-2203311120332113-1110330020220232-1213101202321223-2133202200212011-3301130210013120-3000113220233111-2331332203033110"></a>

Type: `"single"`. Computed.

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.

Upstream description:

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-immediate_action_on_service_down_choice": "[\"immediate_action_on_service_down_drop\",\"immediate_action_on_service_down_none\",\"immediate_action_on_service_down_reset\"]"
}
```

<a id="canonical-3111210113300210-3111002323100302-2033113132131011-3212333331021312-0013130110313232-2001312002230211-1031311232020110-1313220201220311"></a>

## Direct properties — immediate_action_on_service_down / 000132012120 / 3

- [immediate_action_on_service_down_drop](data-sources--application_profiles--reference--group-003.md#canonical-0330110231211133-1013300112302232-0032121321233120-0110023110200023-2011011311313223-3303032311213000-0323031123013102-0310120212202120): complete subsection reference.

- [immediate_action_on_service_down_none](data-sources--application_profiles--reference--group-003.md#canonical-0320222200030133-1100011111332202-3023021021122020-2220000330011332-3001111211301001-3013213030203230-2101133123011203-2323110021001012): complete subsection reference.

- [immediate_action_on_service_down_reset](data-sources--application_profiles--reference--group-003.md#canonical-0223021032303220-1102232003130310-1302011323122020-3231023111301112-3210023001202313-0320211021220022-3320013113301310-2323013330213133): complete subsection reference.

<a id="canonical-0211123031230211-3312231030331020-3102112123311312-2310303123012320-1202121233332010-0013112111020011-1321132032011120-3032033301123222"></a>

## Next pages — immediate_action_on_service_down / 000132012120 / 4

- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](data-sources--application_profiles--reference--group-003.md#canonical-0330110231211133-1013300112302232-0032121321233120-0110023110200023-2011011311313223-3303032311213000-0323031123013102-0310120212202120)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](data-sources--application_profiles--reference--group-003.md#canonical-0320222200030133-1100011111332202-3023021021122020-2220000330011332-3001111211301001-3013213030203230-2101133123011203-2323110021001012)
- [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](data-sources--application_profiles--reference--group-003.md#canonical-0223021032303220-1102232003130310-1302011323122020-3231023111301112-3210023001202313-0320211021220022-3320013113301310-2323013330213133)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0330110231211133-1013300112302232-0032121321233120-0110023110200023-2011011311313223-3303032311213000-0323031123013102-0310120212202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223110000023112-2123330333110002-0022003202313333-0310332033210033-1333302110133010-1130201333033002-3200022111232330-3010211220101232"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop — immediate_action_on_service_down_drop / 233221132210 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop

<a id="canonical-3233003231010201-1213111003322323-2221030030111111-3203212212203322-2322001311323312-0200002223310232-2310212013022231-1232232032330123"></a>

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

<a id="canonical-3310313100112323-2103000120102303-0023100123202330-2112022222001030-3211200031031230-1230020321010022-3033221220021103-1232021203013032"></a>

## Direct properties — immediate_action_on_service_down_drop / 233221132210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311303302133200-3303030213103213-0002120323112330-1322133131321100-2133231003311110-2013112113332233-1320320222000130-3322102121122222"></a>

## Next pages — immediate_action_on_service_down_drop / 233221132210 / 4

- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0320222200030133-1100011111332202-3023021021122020-2220000330011332-3001111211301001-3013213030203230-2101133123011203-2323110021001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113211332021113-3230321301122332-1030110310310112-2303222330121323-2231133003212130-2321101100323021-3003110103203222-3010000011002321"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none — immediate_action_on_service_down_none / 111113222312 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none

<a id="canonical-3000331313012003-1313113120221112-1321320333220223-3223310230103233-3322133303302333-0133332030023102-1120203222130030-3333112121013130"></a>

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

<a id="canonical-0310230112011330-3201333111313120-1110121030320112-2330313232203203-3022012210302320-2323022000223202-2221313011122300-2030133322122023"></a>

## Direct properties — immediate_action_on_service_down_none / 111113222312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310333132210030-1203231002303002-1030202013123312-2003022030131103-1332022011013211-0320030320231110-3201210310310132-0202010330333023"></a>

## Next pages — immediate_action_on_service_down_none / 111113222312 / 4

- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0223021032303220-1102232003130310-1302011323122020-3231023111301112-3210023001202313-0320211021220022-3320013113301310-2323013330213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123100333002222-0301101331321010-2232023001203122-0213003333131101-2332112003013333-0112231222312030-3222123223013222-2110032333031000"></a>

## virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset — immediate_action_on_service_down_reset / 300111223122 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

<a id="canonical-3031312321022010-3333033133022131-3100131033230303-0330010323330111-3111100013013201-0013013212012130-1313113331031321-3002313232001233"></a>

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

<a id="canonical-2233021121233120-2310223231102303-2111111011230202-1320201223102232-1231222023330001-0022211101130202-1202023201033331-1332020333223130"></a>

## Direct properties — immediate_action_on_service_down_reset / 300111223122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202000203303211-2233000030233013-2313021011132003-0233302313300330-3102323011002232-0213120003222333-1301033021211022-0233220300130210"></a>

## Next pages — immediate_action_on_service_down_reset / 300111223122 / 4

- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0000301021323203-1003301002002311-0310011213320313-0102003200331211-3103021001233200-3212113122100001-3320222320102111-3313223020133023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030023320201230-1002302213021210-0300321010113013-0232131103032301-2312031010203322-2320323120213213-1132320110211001-1221023100023133"></a>

## virtual_server.last_hop_pool — last_hop_pool / 202311022302 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.last_hop_pool

<a id="canonical-1331332331330133-1331300103320202-1223101301311323-3321312330210100-2311310201321032-0223302031132302-3110001232223011-0133002310032201"></a>

Type: `"list"`. Computed.

Directs reply traffic to the last hop router using the specified pool.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1203002103030030-2100030121131023-3003111132331322-0110131131030100-0013310230110320-1322321133013201-1312211322021332-3010113300332222"></a>

## Direct properties — last_hop_pool / 202311022302 / 3

<a id="canonical-3320230021230312-1112302231000133-3132223210220133-1333213323222103-2023332210021111-0103333110320202-2121301203201323-1310211133223332"></a>

<a id="canonical-2212221101121321-2001300131122310-0012313002320213-3332321031032030-3222200320013022-1112132312231211-1100230203122233-3030303213300213"></a>

## kind property — last_hop_pool / 202311022302 / 4

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

<a id="canonical-1100132102022032-0321130101023003-1222000130122110-2213212022000021-3122321123032201-0033210123000313-0322011202122302-0203100331021230"></a>

<a id="canonical-1123220103231031-3220032021003312-0330032022010200-3233031330321232-2022221120200112-1221210300132333-1121103013200031-3222003020121323"></a>

## name property — last_hop_pool / 202311022302 / 5

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

<a id="canonical-1022223331012211-3012300010010122-1133012302110113-2111323031231331-3333221002320111-1131010121320023-2223003313103221-3213220311111132"></a>

<a id="canonical-3230021113123010-0203020332220112-3303101333203100-1222211031200110-2013201031002022-1000232320221123-2200121000132131-3022112100211302"></a>

## namespace property — last_hop_pool / 202311022302 / 6

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
  }
}
```

<a id="canonical-0322032110112232-1120333001010313-2222111031233223-3300212301332020-0232023133213333-0010200330202303-0022102030021102-3312021210321200"></a>

<a id="canonical-2021323110200221-3302231123320133-3201102203320122-3123220301312323-3322332201123000-0301022000121231-0231223130221112-1231313201323022"></a>

## tenant property — last_hop_pool / 202311022302 / 7

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

<a id="canonical-0222231132230310-3110230303320121-1133232113122013-1301213010002020-3310012202021002-1233230323311331-2202231211123133-2030012323232211"></a>

<a id="canonical-1321232101102102-3213030111120100-2001331322130303-0122031233232133-3320033302311222-2111222130211000-1002231213302320-3130231012210102"></a>

## uid property — last_hop_pool / 202311022302 / 8

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

<a id="canonical-1100211203301212-1313233113111302-3202313212021310-0102300031332023-3313200230322332-2233123202010222-0300320330013220-3122132132223303"></a>

## Next pages — last_hop_pool / 202311022302 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312232211012033-1001322312123320-0101330122212001-0033002010213211-3202333221111122-1231023111001113-3000133331002211-3123221001232230"></a>

## virtual_server.nat64 — nat64 / 202213112230 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.nat64

<a id="canonical-3112222212333220-2111210012320301-1230013102331102-1113102012213010-1223310022222321-0001023002031323-0231201223233030-0321201011113133"></a>

Type: `"single"`. Computed.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-nat64_choice": "[\"nat64_disable\",\"nat64_enable\"]"
}
```

<a id="canonical-1112201333101101-3323103100202110-1101130001330110-0113330301300321-3113010303212321-3220331330311313-0320000320001033-1331130000313013"></a>

## Direct properties — nat64 / 202213112230 / 3

- [nat64_disable](data-sources--application_profiles--reference--group-003.md#canonical-1022331310002101-0221122112312111-2032231100313013-1022031020210202-3312201210321013-1211121020302023-3201110001133330-3003133303132200): complete subsection reference.

- [nat64_enable](data-sources--application_profiles--reference--group-003.md#canonical-2013103112220012-3012000211330010-0010213022332111-0130021120033031-2221020031113033-2333133200123101-0203122230130021-2130102133201003): complete subsection reference.

<a id="canonical-2113122112300023-0331331321301211-3212122222033221-2322011131232323-0301233110030300-0310133011310121-1102111221012330-0322133231303130"></a>

## Next pages — nat64 / 202213112230 / 4

- [virtual_server.nat64.nat64_disable](data-sources--application_profiles--reference--group-003.md#canonical-1022331310002101-0221122112312111-2032231100313013-1022031020210202-3312201210321013-1211121020302023-3201110001133330-3003133303132200)
- [virtual_server.nat64.nat64_enable](data-sources--application_profiles--reference--group-003.md#canonical-2013103112220012-3012000211330010-0010213022332111-0130021120033031-2221020031113033-2333133200123101-0203122230130021-2130102133201003)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1022331310002101-0221122112312111-2032231100313013-1022031020210202-3312201210321013-1211121020302023-3201110001133330-3003133303132200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121200210122023-0010303011322211-0030012212122032-2200303022102023-3323302031020223-3022023330032321-0010233020120133-3100212312323321"></a>

## virtual_server.nat64.nat64_disable — nat64_disable / 310111121210 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122)
- virtual_server.nat64.nat64_disable

<a id="canonical-3330021000232332-3112120221301132-1202112021311233-1002102220223223-1323103112322321-3020330031020301-1010001112301333-3210122112321303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nat64 disable.

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

<a id="canonical-2112002311230201-1201203200313310-2231122032031121-0012313303100120-1201031103031332-0301101331312321-2220001220232202-0033023320201031"></a>

## Direct properties — nat64_disable / 310111121210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222031321113110-2111120001123311-3301031031111303-0030000312003021-1203112102212002-3123033122231023-3321233022201001-3120000210223323"></a>

## Next pages — nat64_disable / 310111121210 / 4

- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2013103112220012-3012000211330010-0010213022332111-0130021120033031-2221020031113033-2333133200123101-0203122230130021-2130102133201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131022330020201-3023022322001312-2031203111012212-1300021123021210-0123003220333230-2113032010203013-3330120012203112-1233132313210001"></a>

## virtual_server.nat64.nat64_enable — nat64_enable / 232311111232 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122)
- virtual_server.nat64.nat64_enable

<a id="canonical-2323330300333113-0111210302021122-3222102210220312-3120003231033321-1223002012201122-2212212232011330-0022023201212313-2213003331011323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nat64 enable.

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

<a id="canonical-0013022221032121-3321231120123323-1033322022123220-0212203023003321-2200202313131121-3211312311001310-3200300222202131-1101203032210331"></a>

## Direct properties — nat64_enable / 232311111232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002210312230132-0310333032130111-0120112323213110-3202130221123002-1333330321001201-1100310201211302-3100200113131013-2302321121022102"></a>

## Next pages — nat64_enable / 232311111232 / 4

- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112213122012301-2030313331310310-1202112101230331-3213133201121031-2323131123132022-3020020213233300-1023310003202031-0132023003010220"></a>

## virtual_server.port_translation — port_translation / 312310232330 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.port_translation

<a id="canonical-0230012221221331-3021110120120210-2301323100032210-0112002110033330-3020100332131003-3320320221322132-3320012303020331-0002233212311013"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance..

Upstream description:

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

<a id="canonical-2023132220020323-2323130322212311-2323013012330203-0020002322101333-2221133212331232-1020232213002012-3310120010230123-3210220321212123"></a>

## Direct properties — port_translation / 312310232330 / 3

- [port_translation_disable](data-sources--application_profiles--reference--group-003.md#canonical-0301002120332312-1203132031122302-3232230312232323-3302200330030232-1012113113033111-3222222100012201-1023031023301300-0312121120332223): complete subsection reference.

- [port_translation_enable](data-sources--application_profiles--reference--group-003.md#canonical-2133331011212010-0333112201230001-3033230312321133-1100122203213031-2132221222002332-0120313203223322-2221220003310211-3231122203030101): complete subsection reference.

<a id="canonical-0102133232212202-3000331123230222-0331010320202212-1301332230132302-0321220100231010-1320032231211112-3030302120220212-1133030113130132"></a>

## Next pages — port_translation / 312310232330 / 4

- [virtual_server.port_translation.port_translation_disable](data-sources--application_profiles--reference--group-003.md#canonical-0301002120332312-1203132031122302-3232230312232323-3302200330030232-1012113113033111-3222222100012201-1023031023301300-0312121120332223)
- [virtual_server.port_translation.port_translation_enable](data-sources--application_profiles--reference--group-003.md#canonical-2133331011212010-0333112201230001-3033230312321133-1100122203213031-2132221222002332-0120313203223322-2221220003310211-3231122203030101)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0301002120332312-1203132031122302-3232230312232323-3302200330030232-1012113113033111-3222222100012201-1023031023301300-0312121120332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012230132311130-3303002020303232-3112223023220302-0301202221003133-2221203031213003-3013332321012232-2220133312330122-3020200202130211"></a>

## virtual_server.port_translation.port_translation_disable — port_translation_disable / 201203300032 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220)
- virtual_server.port_translation.port_translation_disable

<a id="canonical-1223333011230223-3003313223102211-3213332123113011-2011113321311333-1131201300001232-3222010100131121-3323120023222032-1210311122222011"></a>

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

<a id="canonical-0211302013223021-0302203203032220-3203332112001000-3323021202332030-2213323130011310-0211002203310311-3222222220011220-3323220233221231"></a>

## Direct properties — port_translation_disable / 201203300032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212211330213301-1123300233212122-0223302211230112-2001310203333231-3211111310133320-3220123211212213-2333132110213332-2213231003313022"></a>

## Next pages — port_translation_disable / 201203300032 / 4

- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2133331011212010-0333112201230001-3033230312321133-1100122203213031-2132221222002332-0120313203223322-2221220003310211-3231122203030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000002032112110-3122222331030232-2233233131020313-0011302300333301-0303131313233222-1002120120212233-1203120212322011-1010132232302230"></a>

## virtual_server.port_translation.port_translation_enable — port_translation_enable / 210032223301 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220)
- virtual_server.port_translation.port_translation_enable

<a id="canonical-1123021020332223-3003002223232331-2000213010302201-0232230211330103-1311212121233320-2022302310202033-1333022301001012-0122310131120012"></a>

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

<a id="canonical-0312322231100223-2001110230121120-1201310323020113-1232312023310210-3311201321131230-0313011121030100-1033212023320320-0113121020322301"></a>

## Direct properties — port_translation_enable / 210032223301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333203033111222-2011202331230321-2020221030020230-2222032012013223-0223203001020110-3212032021321312-2000103320310012-3103023321322021"></a>

## Next pages — port_translation_enable / 210032223301 / 4

- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-3232320301212213-2201233011223222-1121012211011023-1212323111231033-1300332330213333-1302020033323222-1223102122223230-1103311321323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303103030002221-3003031100220013-2130033012330213-1012201013020210-3333100231233011-0022133331213300-2232012231231232-3202031221133022"></a>

## virtual_server.request_logging_profile — request_logging_profile / 313020220012 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.request_logging_profile

<a id="canonical-1031110300301322-3011012111331131-3320322013310021-1331332303130013-0333231103331122-3012333020301221-1133212222013310-2220232130012002"></a>

Type: `"list"`. Computed.

Configuration parameter for request logging profile.

Upstream description:

Configuration parameter for request logging profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312211000000102-2221323220003130-3113312211031011-1133333232123013-1010211313203333-0003213210233300-2132112101122333-2100221222002000"></a>

## Direct properties — request_logging_profile / 313020220012 / 3

<a id="canonical-3111011233010123-1322230210312302-3110220021311312-0123010212232222-1221021223020311-2233320201303112-3213221001323133-2222331001120331"></a>

<a id="canonical-3112212121110112-2012313122300300-1113002220222020-0311223332330110-1111103121111132-2113313232110021-1222013311231121-1003333011233030"></a>

## kind property — request_logging_profile / 313020220012 / 4

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

<a id="canonical-0032131021213201-0032210012001322-1022033201203132-0221223233220321-1230311221001013-3311132303202130-1002223322112110-0211321030200212"></a>

<a id="canonical-1120332101230301-2231132301200121-0030320001123010-3220101331223221-1310223112110222-1233203221223001-0123101311132221-2101220111001030"></a>

## name property — request_logging_profile / 313020220012 / 5

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

<a id="canonical-2333231231202222-3000120200101013-2333122312302300-1010300001320031-1312000333330232-3020130310023312-2300030113133222-1332021211122210"></a>

<a id="canonical-1000200211132020-2133122022320230-0100032131032230-3300010310023131-0220210321230210-3121200221012312-3113030233031230-1202111233132131"></a>

## namespace property — request_logging_profile / 313020220012 / 6

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
  }
}
```

<a id="canonical-2321321001100300-2222002201310132-0312100012233022-1110131130102021-1202103112220223-1310120111103121-1312203321021203-0323333131013303"></a>

<a id="canonical-2232211102213301-3122301210123020-3111203203321202-1022123103111000-0113323220232212-2300032032032110-3302220200201333-0111101212101222"></a>

## tenant property — request_logging_profile / 313020220012 / 7

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

<a id="canonical-1211021330212112-0113010233101110-3213130122212220-3122221003303033-3220003021232020-0230220003222200-0120231211301201-3100113311233200"></a>

<a id="canonical-1220132230211311-1003122012023110-2223111100201230-2031231003031232-0201231002113133-2013120230313121-0211210001111312-2122030200113330"></a>

## uid property — request_logging_profile / 313020220012 / 8

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

<a id="canonical-3102111022211112-0123013231101001-3333230000213000-3121133102220113-1232321022020213-0202212223133223-1032022211301020-1303033123301032"></a>

## Next pages — request_logging_profile / 313020220012 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030013022030230-1111202330120203-0213102313311010-0033113230031202-2230022213311211-1211003121021310-2022201211000113-1000120033010202"></a>

## virtual_server.source_port — source_port / 003112123213 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.source_port

<a id="canonical-2231031203231101-1320022022331310-2302331331320012-0013323110030013-0322230203211311-0231131003221110-2123031233033330-3020101331100032"></a>

Type: `"single"`. Computed.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

<a id="canonical-2111332020112220-1130103022003013-1322210131322002-2331002302230121-1030021132133101-1121303303100030-3103211032023230-3303313000203320"></a>

## Direct properties — source_port / 003112123213 / 3

- [source_port_change](data-sources--application_profiles--reference--group-003.md#canonical-1323123103122011-2030010212212212-3032020203020021-0312023101310131-0132200230200033-2130313320131300-3113332320113233-2220123122222023): complete subsection reference.

- [source_port_preserve](data-sources--application_profiles--reference--group-003.md#canonical-0330311022021121-0102221001112332-1012113013030313-1020320310333013-1123321003312213-3233332220100202-0120032020223221-0221012001023101): complete subsection reference.

- [source_port_preserve_strict](data-sources--application_profiles--reference--group-003.md#canonical-1122102321003221-3211213201020003-0010031201013323-2313110220301200-1300131012122331-0030322133003203-0011020313220021-0302002233210021): complete subsection reference.

<a id="canonical-2311112030201321-2023120012200321-1133120231320121-0321233123311130-1321320330000323-2302330033223023-3120223220020032-2331120202332201"></a>

## Next pages — source_port / 003112123213 / 4

- [virtual_server.source_port.source_port_change](data-sources--application_profiles--reference--group-003.md#canonical-1323123103122011-2030010212212212-3032020203020021-0312023101310131-0132200230200033-2130313320131300-3113332320113233-2220123122222023)
- [virtual_server.source_port.source_port_preserve](data-sources--application_profiles--reference--group-003.md#canonical-0330311022021121-0102221001112332-1012113013030313-1020320310333013-1123321003312213-3233332220100202-0120032020223221-0221012001023101)
- [virtual_server.source_port.source_port_preserve_strict](data-sources--application_profiles--reference--group-003.md#canonical-1122102321003221-3211213201020003-0010031201013323-2313110220301200-1300131012122331-0030322133003203-0011020313220021-0302002233210021)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1323123103122011-2030010212212212-3032020203020021-0312023101310131-0132200230200033-2130313320131300-3113332320113233-2220123122222023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301023113231132-0101220122120301-0022322331010022-3222323332300300-2333321202303212-3101032111220131-0010231023312313-1130101000122023"></a>

## virtual_server.source_port.source_port_change — source_port_change / 032023002322 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- virtual_server.source_port.source_port_change

<a id="canonical-1033230100021110-1232321310030300-0013311113121012-3131323201213223-3131121020222213-2123022303233030-0200332113031001-0112011010112101"></a>

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

<a id="canonical-1301322220323300-0122210031220100-3101113311103220-2213022310321321-1120202230220222-2121200330331130-2113113320122310-2103031011311203"></a>

## Direct properties — source_port_change / 032023002322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002311301002333-3112130032222300-1132223323323332-2130322233311311-1323310123312203-2031310002033211-1300130003211023-1132022322231222"></a>

## Next pages — source_port_change / 032023002322 / 4

- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0330311022021121-0102221001112332-1012113013030313-1020320310333013-1123321003312213-3233332220100202-0120032020223221-0221012001023101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323320023102331-2321113030332201-0133001031032103-1231111110103133-0210121302202330-0103320001002011-3322111022330031-0003330320232102"></a>

## virtual_server.source_port.source_port_preserve — source_port_preserve / 313222003312 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- virtual_server.source_port.source_port_preserve

<a id="canonical-3122210132111330-3213033313322121-2230323132323233-0202202102221001-2303103302013330-1203220001233003-0133322023132333-3231231023120323"></a>

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

<a id="canonical-3001220322123213-2133332331323032-1331131110320222-2211300211120033-0032302113230201-0322300010331002-0031030100223132-0021133032011220"></a>

## Direct properties — source_port_preserve / 313222003312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332310121112020-3302112320122303-0123303113310110-2211020112101201-3221100300023311-0033302210202200-0200101323032131-0022031012233101"></a>

## Next pages — source_port_preserve / 313222003312 / 4

- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1122102321003221-3211213201020003-0010031201013323-2313110220301200-1300131012122331-0030322133003203-0011020313220021-0302002233210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331333333230211-2003010232020302-0002221203221333-1123130103221311-2112231313111331-0203312111031202-2001122103133021-3030331233203012"></a>

## virtual_server.source_port.source_port_preserve_strict — source_port_preserve_strict / 322031331103 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- virtual_server.source_port.source_port_preserve_strict

<a id="canonical-2103231131103021-2223110201303233-3123113133113121-2130101121221022-2030110102132300-0102210010213033-2201112112020000-0330100113133120"></a>

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

<a id="canonical-3020301100021111-2330000231230101-0201030113303010-1233132023031100-3131031302022030-3111100332300101-2220332333033133-2212111302321000"></a>

## Direct properties — source_port_preserve_strict / 322031331103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130311213032212-1302223301222301-2301213023000231-0130202021213022-2222313310033030-1002113230233202-1201020032111222-3302122103110131"></a>

## Next pages — source_port_preserve_strict / 322031331103 / 4

- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0230033123202003-3221010301120301-1321302111033133-3300013133310233-0333313030302000-0200200313230110-2223320101301232-1322122301001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203011212210010-2313010312133211-2202021011203003-2320110013013112-2120031123003122-0130321200111203-0232010110310232-3103031223201022"></a>

## virtual_server.statistics_profile — statistics_profile / 123222130021 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.statistics_profile

<a id="canonical-2030120322312202-1311203300012321-0321320031013201-3233002122323031-2033010022232303-2110113120103221-1303322002232120-0232013212222013"></a>

Type: `"list"`. Computed.

Configuration parameter for statistics profile.

Upstream description:

Configuration parameter for statistics profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3210110331310221-1032203031003113-2332132321001123-3010020231333013-0130021201302230-3022010103321211-2220100333301232-3033110322012131"></a>

## Direct properties — statistics_profile / 123222130021 / 3

<a id="canonical-2302010102032211-2101220322000120-1122102130000030-0331230231012103-2033003332311120-1211011002101230-2122113311020032-3233202230133301"></a>

<a id="canonical-2233323200230031-1210000103202131-0130330231202303-3031331000330222-0111110301123303-3323233122110331-3200322132001121-0010301013131020"></a>

## kind property — statistics_profile / 123222130021 / 4

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

<a id="canonical-1210213022223212-3133101111210020-3322303133122300-0123123013113330-2003233011230003-3122232322221131-1113313133211010-2201303221233131"></a>

<a id="canonical-0021301231322122-0032301323112330-2233112331000231-2223123011103132-3122331033210122-3232103101320333-1112323210333012-1112120222230332"></a>

## name property — statistics_profile / 123222130021 / 5

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

<a id="canonical-3233300113322310-1133113133232223-1230022323013133-1211132212130013-2113321221200310-3121033032102200-0220012013130010-0221021211333210"></a>

<a id="canonical-3033221013013321-2220323103030311-0203321102112101-1100033002231301-1303203201001201-1130301111332103-2221333302001120-1202221321333310"></a>

## namespace property — statistics_profile / 123222130021 / 6

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
  }
}
```

<a id="canonical-3200032003300132-3303302103202001-2000030303131112-3100100220110121-3330100231021011-2333200103022023-0302221203020032-2031313331032301"></a>

<a id="canonical-3213101332213203-1321223221013322-0110210320011211-2323320101102113-2111021032111202-3132021033211101-0200300331331223-3001133201311101"></a>

## tenant property — statistics_profile / 123222130021 / 7

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

<a id="canonical-3101021230321013-1312011302130111-3030332132013102-0111212322122021-3203222232012013-1112302322213300-3330331220232323-2233030123302330"></a>

<a id="canonical-1030330003233232-0323020100100002-1223012033010111-3133032011333331-2030211001003120-0022032100032222-3230123303002011-3233300032333210"></a>

## uid property — statistics_profile / 123222130021 / 8

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

<a id="canonical-3232032023303211-2332230222010000-3233021322320202-3032033030021222-2311112112201010-2211321201121303-1101103103311021-3333132102112101"></a>

## Next pages — statistics_profile / 123222130021 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200332232302003-0202220332200232-3002022113000120-0303023020010233-1333321303012321-0221012013330232-1010323021102130-3223110122202211"></a>

## virtual_server.tcp — tcp / 321010222322 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.tcp

<a id="canonical-3332310323031233-0313321330000232-3300010012123032-3321003311001223-1002000031011012-1221013001120331-2103101021011301-1010110123323100"></a>

Type: `"single"`. Computed.

TCP profiles.

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

<a id="canonical-0022332032133122-0000133323333023-2001310212023121-3322320130302020-2231230310013120-3010001113302100-0330202231101132-3211330130323123"></a>

## Direct properties — tcp / 321010222322 / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-0310001233311003-1033213103022230-2312022320203123-2031320231001100-0123321231030330-1010202003122200-1223230132221203-3220003013002122): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-2213210301312212-3113123320032331-3111321222210312-2111022110101033-2213330323231020-3013133312232221-2021112223011032-0302110312212033): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3132332111033112-3021230331103102-0000203200320323-3100311302321003-3203210332032103-3003230230222013-3033030213121211-1000123332303001): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1121221121121022-3213323002333121-1322111332101300-0322022230313330-3212022232023001-2113120221231021-2013111100331123-0231030200223111): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-1231122033130320-3112100322132110-1122202030232302-3121120110303221-2321120012312233-0021322102110131-1211212012311111-2031032102321131): complete subsection reference.

<a id="canonical-1300101020031302-2010031301030221-3132111303010013-0231102200121323-0300132230022213-3311303200031303-2222102222100310-0130020323331023"></a>

## Next pages — tcp / 321010222322 / 4

- [virtual_server.tcp.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-0310001233311003-1033213103022230-2312022320203123-2031320231001100-0123321231030330-1010202003122200-1223230132221203-3220003013002122)
- [virtual_server.tcp.ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-2213210301312212-3113123320032331-3111321222210312-2111022110101033-2213330323231020-3013133312232221-2021112223011032-0302110312212033)
- [virtual_server.tcp.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3132332111033112-3021230331103102-0000203200320323-3100311302321003-3203210332032103-3003230230222013-3033030213121211-1000123332303001)
- [virtual_server.tcp.tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1121221121121022-3213323002333121-1322111332101300-0322022230313330-3212022232023001-2113120221231021-2013111100331123-0231030200223111)
- [virtual_server.tcp.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-1231122033130320-3112100322132110-1122202030232302-3121120110303221-2321120012312233-0021322102110131-1211212012311111-2031032102321131)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0310001233311003-1033213103022230-2312022320203123-2031320231001100-0123321231030330-1010202003122200-1223230132221203-3220003013002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333233011120133-3022112332301232-3013102103020320-0012210320110212-3020031123203203-1033001120311110-2122311003213233-2020000011111322"></a>

## virtual_server.tcp.client_ssl_profile — client_ssl_profile / 022012320320 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.client_ssl_profile

<a id="canonical-1223323031103032-1333301010223312-0323313113233102-2211030313223003-2112002133300232-3131120003201031-3300020231132200-2200333311102021"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2123121121212131-2201303231210031-0113120031231101-3022220022311201-2312001332310212-2101312331103033-1000203232000103-2030022112303123"></a>

## Direct properties — client_ssl_profile / 022012320320 / 3

<a id="canonical-1220010130112121-3220233320100101-3013123332101221-3221133232033011-0022001132301033-2210111002022001-1312132313210102-3332110020230311"></a>

<a id="canonical-3022203011022012-0123301232003220-1121130100320320-2133003212120332-0121022001233012-1033011203101101-1020301033012130-0113210311131130"></a>

## kind property — client_ssl_profile / 022012320320 / 4

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

<a id="canonical-1201232331100312-2231133323321300-1101222002301001-0300120320023002-3312310012202030-2312332002222021-2213223111232022-1122310001101202"></a>

<a id="canonical-3010331202123211-1131223301221101-2100032333130212-0321011202023330-2223232311310330-0032110122021203-0133302220222113-0022300333233010"></a>

## name property — client_ssl_profile / 022012320320 / 5

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

<a id="canonical-0122021320113200-3302220020122013-3130330102311221-3032220300022223-2333010320122220-2033103132200010-0323310333210302-2213121032322330"></a>

<a id="canonical-0102132332212001-0311333212020230-1333321020202223-1311202111121221-1231200030000133-3303110312100033-3012203202320012-2301002123301000"></a>

## namespace property — client_ssl_profile / 022012320320 / 6

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
  }
}
```

<a id="canonical-1012013322002021-2211211113222003-2032221301132030-3223122332033031-3311002131012113-1300301110330220-1032312100101300-0001322103310122"></a>

<a id="canonical-1221133011031233-1330103331332221-1300220011123031-2112223331332011-3222310120232233-0221323313022221-1003310112310002-3011101002021201"></a>

## tenant property — client_ssl_profile / 022012320320 / 7

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

<a id="canonical-2121012201112312-0002001312031121-3201331313001310-2031032313001331-2023333220313133-2013320223201000-2101002012103321-0323010233232122"></a>

<a id="canonical-0222331231131131-0230233221023130-0011211233000332-2221132120002021-2120013202320113-2313030312231223-1312100211000313-1131312031112213"></a>

## uid property — client_ssl_profile / 022012320320 / 8

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

<a id="canonical-3231023233003300-3220223212302011-3210033313321033-0113233120021223-1111131030111031-0113331231222330-2202101231121211-2033332132032303"></a>

## Next pages — client_ssl_profile / 022012320320 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2213210301312212-3113123320032331-3111321222210312-2111022110101033-2213330323231020-3013133312232221-2021112223011032-0302110312212033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320013023203231-2210122121133012-1133020100223022-1210333320310310-3001102021322220-3001103120102032-3001100212301133-3203332330013110"></a>

## virtual_server.tcp.ocsp_profile — ocsp_profile / 122110310312 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.ocsp_profile

<a id="canonical-3222113202322322-0331222010032111-0231200332330131-1120010211000200-1320103023320021-3020032311122013-3012321300232301-1212233311202030"></a>

Type: `"list"`. Computed.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1000121100020120-3330302102031003-0211310200123112-1020002223003320-2220332231001233-1111003003131333-2223320121111012-1232103212213203"></a>

## Direct properties — ocsp_profile / 122110310312 / 3

<a id="canonical-2133103031213311-1023200220030223-0300213312300320-3033203323222030-3003022112303203-3322032133130302-0203212000000210-3210320201220332"></a>

<a id="canonical-1120000001102131-0230100022101302-1201103222130223-1001112102212330-0310320333222011-3031232113111130-3132330013232222-3223132303211123"></a>

## kind property — ocsp_profile / 122110310312 / 4

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

<a id="canonical-1303302102302030-0100000232300311-2210102233131213-3023320231311100-0230322133121211-0012123230110220-1211021233111031-3332133132300322"></a>

<a id="canonical-2021120221212021-0101310120221110-2031312112112021-0103301011103203-2312322011030201-0003021310331321-1112313110030020-3001220120033011"></a>

## name property — ocsp_profile / 122110310312 / 5

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

<a id="canonical-3130230020320011-3103103201100321-1203031032010003-0133021210001121-0021231011103331-3233220302011033-3212223022020031-2202112331033011"></a>

<a id="canonical-2032203110131313-0032323330300322-1110303221301200-3033030002321000-1032012002011230-2100103213002200-0303332203130312-0120302130003123"></a>

## namespace property — ocsp_profile / 122110310312 / 6

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
  }
}
```

<a id="canonical-0322021303013323-0202320023211102-3100101303010000-2210122222101322-2322233231000131-2303121110023200-1220332322133203-1222233023231233"></a>

<a id="canonical-0131131310030102-2002221312011321-1110111003113220-3300232113201121-2222100300222101-1310113110330230-0122230031131133-0233033310302320"></a>

## tenant property — ocsp_profile / 122110310312 / 7

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

<a id="canonical-3233130311023002-2222211223202010-1211231210210222-1320312111330103-3231201020300111-3033322213122232-0100232132220223-1321323302211202"></a>

<a id="canonical-2123021000223230-2211333022031113-1312201222013220-1331012311310012-0110102000223102-1300311100031231-1131012101111302-1202032023022230"></a>

## uid property — ocsp_profile / 122110310312 / 8

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

<a id="canonical-2203300101121223-1313003000322322-1212103000002033-3212113102213121-0131132102101222-3301221300123103-1002112323321001-2221110001113030"></a>

## Next pages — ocsp_profile / 122110310312 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-3132332111033112-3021230331103102-0000203200320323-3100311302321003-3203210332032103-3003230230222013-3033030213121211-1000123332303001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113000111313130-3303133222102223-0323311300210010-1301133022120032-1131100233013112-0032013112302212-0331022010210030-3331202020203111"></a>

## virtual_server.tcp.server_ssl_profile — server_ssl_profile / 200023332030 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.server_ssl_profile

<a id="canonical-1301210030113010-3001031000022023-1203202300132011-2320223221023320-0201113030211233-1113112203320131-0321103311130200-2313302302001132"></a>

Type: `"list"`. Computed.

Configuration parameter for server SSL profile.

Upstream description:

Configuration parameter for server SSL profile

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2100232120002203-3113211212122031-0100223133133001-2103213223110202-1013101223231302-2011010020312023-3301203123000123-2222330201120322"></a>

## Direct properties — server_ssl_profile / 200023332030 / 3

<a id="canonical-3123220302001003-0221110312110212-3010023023031232-0100021213122233-0231101221102032-3010020223123303-3201231220101332-3233230022333100"></a>

<a id="canonical-0332212233223102-0130032323012002-0112011121100220-2330302213323213-2203023213321103-0230310233222002-3010010130202211-1320100121003330"></a>

## kind property — server_ssl_profile / 200023332030 / 4

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

<a id="canonical-3220023021330010-1022331102300302-1130310302211113-2203331201322310-0233211102323232-1103213012203022-2011230332230113-0132221121220122"></a>

<a id="canonical-1112120111323233-3022312233310010-3010133011321311-2203021312030112-0302003032301201-0122120231110231-2021223221200303-2030211111132331"></a>

## name property — server_ssl_profile / 200023332030 / 5

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

<a id="canonical-3111223102012123-0203023033221220-0332303002133231-2031010210000102-1020011001012231-1102000000113011-3002330230133202-1213201203331323"></a>

<a id="canonical-2013102031300303-2221321202330131-2020231121231211-3013332302000100-3201303003330022-3303003232233011-0323032222321003-2000231020223331"></a>

## namespace property — server_ssl_profile / 200023332030 / 6

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
  }
}
```

<a id="canonical-2022110203301233-2233230031101202-1300223210332300-3321000123200220-1131320332213333-2202311133113200-2133233021110212-0032013030303331"></a>

<a id="canonical-0203200330312013-2023312110233330-2030202111220331-2310331023203303-1203211112012230-0101101220213130-0003200212333012-1322232330112203"></a>

## tenant property — server_ssl_profile / 200023332030 / 7

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

<a id="canonical-0103200032003121-2303201330033021-0032320131330211-3131000310203303-3233202112120102-1231321133321301-3102221230101311-1003122022310022"></a>

<a id="canonical-2021123223230222-2012331331221201-3310310220201222-3221230121311320-2201022113211032-1332203013133031-3230031133020332-2211013312313030"></a>

## uid property — server_ssl_profile / 200023332030 / 8

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

<a id="canonical-2032211131322332-0321012213121312-1113012213202131-1200030020010010-0031212013332333-0303132123011123-2032220123210121-0221200010321223"></a>

## Next pages — server_ssl_profile / 200023332030 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1121221121121022-3213323002333121-1322111332101300-0322022230313330-3212022232023001-2113120221231021-2013111100331123-0231030200223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210023332310012-2320320030000120-2022232132312122-3331331211000022-3022322223203023-3121321031122303-0332013202220133-2221323112200231"></a>

## virtual_server.tcp.tcp_client_profile — tcp_client_profile / 031230203033 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.tcp_client_profile

<a id="canonical-0233130021010220-1121023203223312-3032230023001021-2100113110302023-2010233012010211-1320211021101331-2021110312323310-1030121232022110"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2002203012221300-1321222022122023-3102230201333021-2210012130320013-1300302002213110-3221031322212202-0323202122111103-1212131231331031"></a>

## Direct properties — tcp_client_profile / 031230203033 / 3

<a id="canonical-3032220312231213-2300232121210230-2212123331131322-1031321232021130-3023223320313113-2021331121331212-1331312010331211-2032022210121023"></a>

<a id="canonical-0113122302112201-3313202013301001-1101202320002200-3322201222112233-0033132320213220-0302310321301123-3312100301311231-3132133020203102"></a>

## kind property — tcp_client_profile / 031230203033 / 4

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

<a id="canonical-0303213032221112-1100121300200200-1320222322310231-3123002211221131-1230311310232123-3131301212131303-1102221213030213-2031300131331031"></a>

<a id="canonical-0320300122122031-0301003211223103-1010021102310201-1101321311322223-2123010110001202-0301032100132033-1212232313221032-0120032232000222"></a>

## name property — tcp_client_profile / 031230203033 / 5

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

<a id="canonical-2001200320213233-1120030103320132-1032112022111031-3211100010100330-2302333212033031-2133231220001010-2233111313032112-3212032333310201"></a>

<a id="canonical-0022000203301301-2310012313332210-0330101003100332-3033330321020121-2201031330203011-0231213030010222-1312131100020320-3313323212030200"></a>

## namespace property — tcp_client_profile / 031230203033 / 6

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
  }
}
```

<a id="canonical-1212100302003313-0331210323332030-1201303123012230-0030221332203201-3001113132301312-3101032331100222-3233223310013113-1110332333101122"></a>

<a id="canonical-1303211023020120-1020002011203203-0330103103233211-3213102132100333-1123200033113322-1303303320300300-3221112123010000-2203302111310310"></a>

## tenant property — tcp_client_profile / 031230203033 / 7

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

<a id="canonical-0233221223033011-1212023121031111-0321033132333302-0003331011310333-1113023121012311-1020210031000302-1221023202000030-0220012210312210"></a>

<a id="canonical-0030022031231001-3101310222332131-1101333002021323-1122012030030313-1310322300002330-2022302210122230-3231032203330121-3223113312212010"></a>

## uid property — tcp_client_profile / 031230203033 / 8

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

<a id="canonical-0200123102111231-3100212020131132-2230033020023223-0131011231223213-0202132210230130-2231200023202031-3032133022220003-0132301301033133"></a>

## Next pages — tcp_client_profile / 031230203033 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1231122033130320-3112100322132110-1122202030232302-3121120110303221-2321120012312233-0021322102110131-1211212012311111-2031032102321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313022020310122-2122312111210020-0022303233213203-2303221010302213-1230112133022333-0331101113012203-0133032121031200-1320122202323310"></a>

## virtual_server.tcp.tcp_server_profile — tcp_server_profile / 320102322211 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.tcp_server_profile

<a id="canonical-2223101012002002-0211221222230013-2203113222222221-2000010023113232-0022101110310031-3312130211212133-3101130033321322-2112213130333221"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103010000310203-2100032001231000-2102311101212301-0303230130233302-2231223030000332-0012320232112131-3100300003213111-0233111011112230"></a>

## Direct properties — tcp_server_profile / 320102322211 / 3

<a id="canonical-2331002323321112-1113033321031011-1321232120031023-1023222002103233-2021021130003312-0301112010210103-0120110110023201-2002121230011020"></a>

<a id="canonical-2110330032020220-3213113011013131-0321231333300213-3133022002111121-2111131322120023-3232210030320001-2001013113222110-3313133130101230"></a>

## kind property — tcp_server_profile / 320102322211 / 4

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

<a id="canonical-2003031013322331-0100100330122133-0303211313113032-2110233133231212-3012313200011313-2031020020322202-1020332111121023-1223121200213100"></a>

<a id="canonical-2013021302131321-3210010103330022-2010320221321210-3333101023221123-2331033030332220-0111202321312113-3122002222001033-3131122222200033"></a>

## name property — tcp_server_profile / 320102322211 / 5

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

<a id="canonical-3322220201010210-0203220112133211-1120123201023000-0231313101033100-1332111013202002-0310230131120331-2321312000120302-2102230102110023"></a>

<a id="canonical-2301302113133312-1311033221033102-2223331000333001-3111132020133300-3122233022301120-3102020232102211-3223031330131103-2001210100121303"></a>

## namespace property — tcp_server_profile / 320102322211 / 6

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
  }
}
```

<a id="canonical-0321123000011210-2010112330021131-1331101313300303-2310210120100000-1323223202001223-1130031103303331-1222032231310200-2101120201110231"></a>

<a id="canonical-0123113122200121-1233020123020011-0000312000030212-0000233030300220-3112011031200110-2231213110130320-2000023033112030-0001213122323021"></a>

## tenant property — tcp_server_profile / 320102322211 / 7

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

<a id="canonical-3112202321320130-0302202312113323-0330303200001100-3323010223012021-2212010221311202-3223032212230232-1231023121312230-0110130011003100"></a>

<a id="canonical-1232232111022013-3032021211100110-1120112130110100-1001131203330220-2313101131123222-0331312202111023-1011310221203110-1122313110112322"></a>

## uid property — tcp_server_profile / 320102322211 / 8

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

<a id="canonical-1111030002032322-0303020231112301-0131320130032120-2110221002311223-1003010022033310-2023330302220111-0133032013203211-2101223211123031"></a>

## Next pages — tcp_server_profile / 320102322211 / 9

- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123102213120331-0211003211120013-1223210122230113-0011130230302213-2313212302033223-2023030002103203-0321313210132103-2321301330330330"></a>

## virtual_server.udp — udp / 003302030230 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.udp

<a id="canonical-3230102230310120-1010023200321221-3133023233101012-0120203133222230-3233210233210103-1332031102321121-2231002312110123-2021321301012232"></a>

Type: `"single"`. Computed.

UDP profiles.

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

<a id="canonical-2001121013100202-1232122203112302-0130311011203310-0330011220000321-2121320013032103-2332111111321003-2020310320331331-2133030130100321"></a>

## Direct properties — udp / 003302030230 / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1130220011130010-2223332122330201-2233130011332112-1222012131322323-2032233002322102-1330122332023010-2310031123300223-3001301201213031): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1233221132021202-2231130020101220-2012310103223310-1012001123132022-2002120023221313-0323030301120010-0302122033300222-0013002113111320): complete subsection reference.

- [udp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1333123302300223-1003212003320120-3011131021320211-3121211002322222-0101202302010313-3232321330230032-2023323301311101-0003010123121122): complete subsection reference.

- [udp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-0223131020311101-3120313230320201-0310233233131121-2012011133331201-0212100200033333-0303313300113123-1231122103120022-2211230303302203): complete subsection reference.

<a id="canonical-0210322313221312-1100333322010131-0032100000013011-1033321333220100-1322121002113112-1330321102331030-1110020213030302-3022120303012000"></a>

## Next pages — udp / 003302030230 / 4

- [virtual_server.udp.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1130220011130010-2223332122330201-2233130011332112-1222012131322323-2032233002322102-1330122332023010-2310031123300223-3001301201213031)
- [virtual_server.udp.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1233221132021202-2231130020101220-2012310103223310-1012001123132022-2002120023221313-0323030301120010-0302122033300222-0013002113111320)
- [virtual_server.udp.udp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1333123302300223-1003212003320120-3011131021320211-3121211002322222-0101202302010313-3232321330230032-2023323301311101-0003010123121122)
- [virtual_server.udp.udp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-0223131020311101-3120313230320201-0310233233131121-2012011133331201-0212100200033333-0303313300113123-1231122103120022-2211230303302203)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1130220011130010-2223332122330201-2233130011332112-1222012131322323-2032233002322102-1330122332023010-2310031123300223-3001301201213031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301002120223123-1131111310032203-0203121333111332-3011100332231011-1012031122121101-2112230312111001-1121212212322131-2320011131021133"></a>

## virtual_server.udp.client_ssl_profile — client_ssl_profile / 010121311312 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.client_ssl_profile

<a id="canonical-2223111031211133-0101333002323221-1330302200031101-2132112223123222-2201122023213030-1020321313300011-1010133331202031-3230300310130112"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0222211331100211-0111032221102221-0213021112131032-2322011133021230-2310013311122200-3000213001131032-0230230102022320-3030322030323211"></a>

## Direct properties — client_ssl_profile / 010121311312 / 3

<a id="canonical-0322301012103310-2211210202321302-2220100202033221-3123231121221030-2222333013100200-2223332112112210-1131021123131111-3231210003232133"></a>

<a id="canonical-3200212111130220-3223231002011210-2010213022133131-0330223300020211-2311000200231001-1230131322232230-1331202311113332-0023113033000200"></a>

## kind property — client_ssl_profile / 010121311312 / 4

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

<a id="canonical-0132103230323301-2310210101023022-3221100322330230-2020223323320212-0212212311121321-0113221020133331-2303003101010200-2122303332230211"></a>

<a id="canonical-0102320313100000-2310023020033120-1302223323202103-1021113221031320-3100312210300122-3012202310000003-0031123030023101-3120003320323010"></a>

## name property — client_ssl_profile / 010121311312 / 5

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

<a id="canonical-1210103233123230-0011200330100300-3122332312110310-0233010331020013-1122002212112101-1212203332320202-1321202211323003-1102301221013010"></a>

<a id="canonical-0021203102322203-3011332213312113-0211320022033223-1122023333333031-3121132102320021-3302331022133133-1103011222310223-2213011002310313"></a>

## namespace property — client_ssl_profile / 010121311312 / 6

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
  }
}
```

<a id="canonical-0211311122032003-3222130111113002-3101000323101010-2021113023000331-2111000102130221-2111120203222011-1100013033011023-0203132122212300"></a>

<a id="canonical-0022201200120322-1220022231120031-2032122011310002-0031322112100033-2002233000231201-0233323132321032-1200202031023301-3100131330320001"></a>

## tenant property — client_ssl_profile / 010121311312 / 7

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

<a id="canonical-2211100301020302-2320012111132033-1201303120022021-1220213100122220-2012021300310132-0233030333313032-2011213112313003-0231302123123110"></a>

<a id="canonical-0011002011311001-3312123100332332-2222231100012200-0102121123320123-1013210103201311-0331322233013330-2110021001001333-0031103012110220"></a>

## uid property — client_ssl_profile / 010121311312 / 8

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

<a id="canonical-2211301103331322-0030313321220221-0310012231303201-2131330201321210-3220321302233103-2220133122001332-0110221232121120-1201331130010330"></a>

## Next pages — client_ssl_profile / 010121311312 / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1233221132021202-2231130020101220-2012310103223310-1012001123132022-2002120023221313-0323030301120010-0302122033300222-0013002113111320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032222111023030-0131112222003030-1011223321001313-3003230222200331-2113112312301210-0231201213103332-0231121023122112-3112031232322201"></a>

## virtual_server.udp.server_ssl_profile — server_ssl_profile / 032000110111 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.server_ssl_profile

<a id="canonical-0101033012301333-2330013230023300-1030200021120003-3111233103113030-1102102100332031-2131121031322222-2213131312112330-2301020232113100"></a>

Type: `"list"`. Computed.

Configuration parameter for server SSL profile.

Upstream description:

Configuration parameter for server SSL profile

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131022310222203-3232022200111110-3122331110231232-1311233121301123-3203200200203312-0321132303113131-1130020010011203-2020131033133303"></a>

## Direct properties — server_ssl_profile / 032000110111 / 3

<a id="canonical-2332323020202202-2301010013113311-3201203231333312-0131221313130303-2301001121121100-0303113122120330-0132303133110002-2020321130031112"></a>

<a id="canonical-0032210201011002-3312100100012123-0330300123100001-0130221120230012-0120210030101331-3320221330210020-2331210002111323-0233233211020033"></a>

## kind property — server_ssl_profile / 032000110111 / 4

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

<a id="canonical-2221303121101223-0133032033012012-3313000223203101-3331221001133121-2010132220010230-0231313311333303-1323002202323311-1301012013102131"></a>

<a id="canonical-0022021230201103-2333133132012010-1133322300321313-2322333330310202-1030220313231021-1012233030121012-2033003130312100-3313200100112023"></a>

## name property — server_ssl_profile / 032000110111 / 5

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

<a id="canonical-1100203111120031-2120323223132010-2301303001222002-3012020030133313-1213312200202111-0210120101322223-2301011312021011-2000111012011011"></a>

<a id="canonical-0010120010023033-2022000221031010-0032012232132100-2202110221111020-2101111131322111-2231303333003023-3331222213020030-2132201233132203"></a>

## namespace property — server_ssl_profile / 032000110111 / 6

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
  }
}
```

<a id="canonical-1223122120321102-2000203333123003-1322122023312212-3000031011123100-3212313310230210-2312232300021021-2013231031020232-1302021022310010"></a>

<a id="canonical-1313033131302023-0210222030211032-3203202100232133-3301232011023000-0112130032203203-2333202200223011-2302220232022030-0203223220232231"></a>

## tenant property — server_ssl_profile / 032000110111 / 7

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

<a id="canonical-2231130232100022-3123001103220031-2033212310210310-2000023222321131-0000302302133013-2212322120103331-0210122331000112-0102011101103103"></a>

<a id="canonical-3031302111120133-3131300331200122-3303001022101002-2233333113302223-0010331011131222-2302303311331321-0120200121013333-2001032113112003"></a>

## uid property — server_ssl_profile / 032000110111 / 8

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

<a id="canonical-3031333311133323-2100111202222020-2112200320222110-3313012212023310-2303022030323312-0210000023210023-0002230012001120-3112012012333132"></a>

## Next pages — server_ssl_profile / 032000110111 / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-1333123302300223-1003212003320120-3011131021320211-3121211002322222-0101202302010313-3232321330230032-2023323301311101-0003010123121122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301210302022132-0020130322110121-3212332231103000-2220010231321101-1121210123212000-3331231200313121-2313212213032213-2111102110313201"></a>

## virtual_server.udp.udp_client_profile — udp_client_profile / 303232000330 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.udp_client_profile

<a id="canonical-1022022013310211-3131023311012011-0332223013233330-0310123313002302-3003030100013303-1331103101100223-0002131121201323-2223000021033100"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320111222213100-3231320333013320-1323221302220200-3100321131210321-0103231030033221-1002333030233011-1231030110210213-0230130031033300"></a>

## Direct properties — udp_client_profile / 303232000330 / 3

<a id="canonical-3331321213300032-2002032122201223-3132221303332332-3112013011102232-2303100021103021-1210120010303210-3203033000020233-0323212333002033"></a>

<a id="canonical-2230010210211302-3022101312021012-0313010120022311-1202121201221210-0313201121001130-0132132131332323-3100022200200113-3003001230213130"></a>

## kind property — udp_client_profile / 303232000330 / 4

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

<a id="canonical-1320223331022233-0211201230233120-0223002213100200-3210120200222330-2133220001131202-2031210132300313-2232233032130111-3320131033220202"></a>

<a id="canonical-3201000211322032-2111322311331122-3323221103332121-3030313023322300-1202110211203313-0333312100103222-1210321132212031-1301102220332223"></a>

## name property — udp_client_profile / 303232000330 / 5

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

<a id="canonical-2303121332220020-2200102003102322-3210233320133012-3021031022032321-2130102310302211-3320012122133032-1213132002202031-3321312330233131"></a>

<a id="canonical-3122130230320101-3203103031110002-0122200103112033-1112133222132323-3331232303100131-3311221320003331-1203020212121003-3030020301312012"></a>

## namespace property — udp_client_profile / 303232000330 / 6

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
  }
}
```

<a id="canonical-1111313003330213-2112303201023030-0332303020331322-0221001122220130-3311311030311033-3231130330213022-3301032100122122-3323200303203223"></a>

<a id="canonical-0232211033303313-0311110131212321-3102100002222131-3312103122112020-2210002323223301-2013123301202310-0100300112223223-3112002210311323"></a>

## tenant property — udp_client_profile / 303232000330 / 7

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

<a id="canonical-3123323202130100-3132220203121313-3210312130122012-1311201213301201-2212033102311122-2000111122023232-1101032320213121-3232013020101233"></a>

<a id="canonical-2122311232323210-0133203111212130-0233130210133223-2011132330213031-0330112302320200-1221100030110231-3002000010010313-3132123123221103"></a>

## uid property — udp_client_profile / 303232000330 / 8

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

<a id="canonical-2013002232000002-2221100331031133-3231021203322323-3220102113232220-2223211203213333-3202301123122232-1211212322230303-0133032310101320"></a>

## Next pages — udp_client_profile / 303232000330 / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0223131020311101-3120313230320201-0310233233131121-2012011133331201-0212100200033333-0303313300113123-1231122103120022-2211230303302203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032330312030011-0031331201200120-0223011323013211-2023313113232321-0223213313011220-3202100000302102-3311203320000323-3130310110123323"></a>

## virtual_server.udp.udp_server_profile — udp_server_profile / 131002131213 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.udp_server_profile

<a id="canonical-1232210232213322-3200033101023210-1022100231121012-3311323303321202-0130032002133020-2301003122033333-3021010230323203-3032301322031133"></a>

Type: `"list"`. Computed.

Configuration parameter for udp server profile.

Upstream description:

Configuration parameter for udp server profile

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3012022333322231-2221313102110200-0312033122212223-1031103012333032-1213111332021213-1111013200112310-2100202201233103-1130011332330333"></a>

## Direct properties — udp_server_profile / 131002131213 / 3

<a id="canonical-0301232220333320-0212023310223321-2233100220312112-2300331223010303-2011223202200333-1133121101223323-2123120033132120-0231132012112322"></a>

<a id="canonical-1013330012330101-3022133131323022-0230231332301333-2313120311102201-0131031311230030-3210233322203202-3322331111221223-1333212130033002"></a>

## kind property — udp_server_profile / 131002131213 / 4

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

<a id="canonical-1301113121300332-2011130120100301-2332131131321130-2110210213232321-3331311222230133-2321020223101121-1033201330333101-3311123323211313"></a>

<a id="canonical-2111321131320132-0230330323312120-2331032221030312-0300022022320122-2221130122211023-2220302002111122-1132030232213111-3301332323210033"></a>

## name property — udp_server_profile / 131002131213 / 5

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

<a id="canonical-1100321230312010-3023132212020002-3332131310320113-0331330002112322-1132021322313012-1211330321022323-0202220332323020-1300313033013323"></a>

<a id="canonical-2301120023020030-1302230102012010-3302022112130110-3301131130121012-3012322301301131-1312200131313331-2123121332131231-1033132022233212"></a>

## namespace property — udp_server_profile / 131002131213 / 6

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
  }
}
```

<a id="canonical-0020233331120310-1320122000011301-1300121301323333-1233233113202320-3223012111320202-0231232222032033-2103203000222311-1012123333332121"></a>

<a id="canonical-3232000322322110-3100313023300302-3220112003102032-3031332313021202-2333131332120230-2011012201113320-0203022101130101-3331322232300332"></a>

## tenant property — udp_server_profile / 131002131213 / 7

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

<a id="canonical-3302132310200220-1112211012111031-0123230123211311-2131012311210121-3310321130120201-2201123213011210-1123332101010032-1222202211310232"></a>

<a id="canonical-1012221103330110-1320033221230320-2212102123001231-0300121001012012-0012322102323331-1122003012101320-2211011123103120-0231023321231333"></a>

## uid property — udp_server_profile / 131002131213 / 8

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

<a id="canonical-3211130333200221-2011320313322212-3111110332321022-2202322013303111-2020113321121102-3001133203122123-0101030102231033-3003332211322103"></a>

## Next pages — udp_server_profile / 131002131213 / 9

- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-0300001203202301-3231100101123312-2010220203023303-2030320120111312-3201010231210333-1312201331313222-1321111320032321-3200132220130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333033230000332-0021310100001130-0021023213233010-3332231111100332-1230032033003120-0003312231011111-0200201331200011-3100010301130332"></a>

## virtual_server.virtual_server_state — virtual_server_state / 003300311331 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.virtual_server_state

<a id="canonical-2021230121300320-1123230031330120-2201000331133031-3132210011103321-1233130023200123-0213223330301201-0132022232032021-1202010313232100"></a>

Type: `"single"`. Computed.

Displays the current state on the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

<a id="canonical-0323112302123110-2033001121213310-2110312333301300-0200323123200231-3133123301330201-3321121002132122-3222231230023131-1210312230321300"></a>

## Direct properties — virtual_server_state / 003300311331 / 3

- [state_disabled](data-sources--application_profiles--reference--group-003.md#canonical-2212232221003212-1302330323113311-3333332222133201-0120022103313000-3321101301200111-3002031100221102-3323101001203133-0111000121101200): complete subsection reference.

- [state_enabled](data-sources--application_profiles--reference--group-004.md#canonical-3203221331230032-2132020121120220-2200010100021202-3030323231022121-2210002223020313-3132333233332103-3321200032123011-2330111003302323): complete subsection reference.

<a id="canonical-3103301233020100-2231311202223320-2230132221331312-1121111231111321-1322331030011323-0002100112000021-2110313232000212-3312023131310212"></a>

## Next pages — virtual_server_state / 003300311331 / 4

- [virtual_server.virtual_server_state.state_disabled](data-sources--application_profiles--reference--group-003.md#canonical-2212232221003212-1302330323113311-3333332222133201-0120022103313000-3321101301200111-3002031100221102-3323101001203133-0111000121101200)
- [virtual_server.virtual_server_state.state_enabled](data-sources--application_profiles--reference--group-004.md#canonical-3203221331230032-2132020121120220-2200010100021202-3030323231022121-2210002223020313-3132333233332103-3321200032123011-2330111003302323)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)

<a id="canonical-2212232221003212-1302330323113311-3333332222133201-0120022103313000-3321101301200111-3002031100221102-3323101001203133-0111000121101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112301231121201-0212202203000233-0311111130223103-0023112300001123-3332022321103221-2033332300331112-1223033332331021-3211122030033032"></a>

## virtual_server.virtual_server_state.state_disabled — state_disabled / 301300020133 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-0300001203202301-3231100101123312-2010220203023303-2030320120111312-3201010231210333-1312201331313222-1321111320032321-3200132220130130)
- virtual_server.virtual_server_state.state_disabled

<a id="canonical-2333230033233103-1022032110001213-1211312103321332-2213010220333300-1321201012201122-0222101313111331-2101220020032332-2311032233010201"></a>

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

<a id="canonical-0323033310302222-1310020013200202-3301101231110203-2001332022133001-2223232231020303-1210303122221312-3100312132131223-2002101123020223"></a>

## Direct properties — state_disabled / 301300020133 / 3

This is an empty object or choice marker. It has no direct properties.
