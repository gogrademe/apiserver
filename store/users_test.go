package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	. "gopkg.in/check.v1"
	// "log"
)

// func (s *RethinkSuite) TestAggregationReduce(c *test.C) {
// 	var response int
// 	query := Expr(arr).Reduce(func(acc, val Term) Term {
// 		return acc.Add(val)
// 	})
// 	res, err := query.Run(sess)
// 	c.Assert(err, test.IsNil)
//
// 	err = res.One(&response)
// 	c.Assert(err, test.IsNil)
// 	c.Assert(response, test.Equals, 45)
// }

// TestUserDatabase verifies that a User can be saved and loaded from the database
func (s *StoreSuite) TestUserStore(c *C) {
	for i, t := range []struct {
		summary    string
		email      string
		password   string
		role       string
		shouldFail bool
		checkErr   error
	}{
		{
			summary:  "valid user",
			email:    "test@test.com",
			password: "somePassword",
			role:     "Admin",
		},
		// {
		// 	password: "somePassword",
		// 	role:     "Admin",
		// },
		// {
		// 	role: "Admin",
		// },
		{
			summary:  "duplicate user",
			email:    "test@test.com",
			password: "somePassword",
			role:     "Admin",
			checkErr: ErrUserAlreadyExists,
		},
	} {
		u, err := m.NewUser(t.email, t.password, t.role)
		err = Users.Store(u)
		c.Logf("test %d: %s", i, t.summary)
		c.Assert(err, Equals, t.checkErr)
	}
}
