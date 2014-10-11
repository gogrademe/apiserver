package store

// // QueryAllPeople return query for all people without their profiles.
// func QueryAllPeople() r.Term {
// 	query := People.OrderBy("firstName", "middleName", "lastName").Map(func(row r.Term) interface{} {
// 		return row.Merge(map[string]interface{}{
// 			"profiles": map[string]interface{}{
// 				"studentId": r.Table("students").Filter(func(s r.Term) r.Term {
// 					return s.Field("personId").Eq(row.Field("id"))
// 				}).CoerceTo("ARRAY").Map(func(s r.Term) interface{} {
// 					return s.Field("id")
// 				}).Nth(0).Default(""),
// 				"teacherId": r.Table("teachers").Filter(func(s r.Term) r.Term {
// 					return s.Field("personId").Eq(row.Field("id"))
// 				}).CoerceTo("ARRAY").Map(func(s r.Term) interface{} {
// 					return s.Field("id")
// 				}).Nth(0).Default(""),
// 			},
// 		})
// 	})
//
// 	return query
// }
