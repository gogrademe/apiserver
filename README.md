GoGrade
=======

r.db("dev_go_grade").table("enrollments").filter({
    classId: "c0c54f06-3189-4794-8dca-5b885f355e5b",
    termId: "08ad28d4-7f5a-4475-8607-347f79526590"
}).eqJoin("studentId", r.db("dev_go_grade").table("students")).map(function (var_30) {
    return var_30("left").merge({
        student: var_30("right")
    })
}).eqJoin(function (var_31) {
    return var_31("student")("personId")
}, r.db("dev_go_grade").table("people")).map(function (var_32) {
    return var_32("left").merge({
        person: var_32("right")
    })
}).orderBy(r.asc(function (var_33) {
    return var_33("person")("firstName")
}), r.asc(function (var_34) {
    return var_34("person")("middleName")
}), r.asc(function (var_35) {
    return var_35("person")("lastName")
}))
