db.createUser({
  user: "tweeto",
  pwd: "secret",
  roles: [
    {
      role: "readWrite",
      db: "tweeto",
    },
  ],
});
