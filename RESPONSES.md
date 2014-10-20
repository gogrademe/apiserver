Every response from the api must follow this format, including errors, searches, gets, posts, authorization, **all responses**.  UPPER_CASE is used to represent a variable.  They're defined following this snippet.

```json
{
    "TYPE": [
      {
        "id": "THING_ID",
        "OTHER_TYPE_ID": "ID",
        "OTHER_TYPE_ID_MULTI": ["ID", "ID"],
        "MISC_PROPERTY": "$MISC_VALUE",
        "OTHER_MISC_PROPERTY": "$OTHER_MISC_VALUE"
      }
    ]
}

```

<dl>
  <dt> `TYPE`
  <dd> type is the resource type.  Examples: person, student, teacher, grade, assignment, session
  <dd> there may be multiple values for a type
  <dd> there may be multiple types on a response
  <dd> `GET /api/TYPE/id` should give the data for an object type

  <dt> `id`
  <dd> an id that is unique to this object within its type
  <dd> `GET /api/TYPE/id` should give the data for an object type

  <dt> `OTHER_TYPE_ID`
  <dd> another type name such as: personId, studentId, teacherId, etc.
  <dd> removing the "Id" suffix is the TYPE of that object; `GET /api/TYPE/ID` should resolve that object

  <dt> `OTHER_TYPE_ID_MULTI`
  <dd> the key is the same as `OTHER_TYPE_ID` e.g. personId, studentId, teacherId
  <dd> if it's possible for there to be multiple related documents of the same type, it should have an array response

  <dt> `MISC_PROPERTY`
  <dd> any property the gives more data about this object, but not any of its related objects

  <dt> `MISC_VALUE`
  <dd> this may be a string, number, boolean, an object or array of objects with `"MISC_PROPERTY": MISC_VALUE` pairs

</dl>
