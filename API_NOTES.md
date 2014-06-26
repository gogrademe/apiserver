
```
  r.db("dev_go_grade").table("people").merge(function(person) {
    return {
      profiles: {
        student: r.db("dev_go_grade").table("students").filter(function(student) {
          return student("personID").eq(person("id"))
        }).coerceTo("ARRAY")
      }
    }
  })
```

```
{
  "createdAt": Mon Jan 01 1 00:00:00 GMT+00:00,
  "firstName":  "Jon",
  "id":  "2dbc1a60-55a7-4762-b28a-ade1333eb84e",
  "lastName":  "Bush",
  "middleName":  "David",
  "profiles": {
    "student": [
      {
        "createdAt": Mon Jan 01 1 00:00:00 GMT+00:00,
        "gradeLevel":  "12th",
        "id":  "fb5a7702-220f-436c-88f5-3cad009e1ea8",
        "personID":  "2dbc1a60-55a7-4762-b28a-ade1333eb84e",
        "updatedAt": Mon Jan 01 1 00:00:00 GMT+00:00
      }
    ]
  },
  "updatedAt": Mon Jan 01 1 00:00:00 GMT+00:00
}
```
