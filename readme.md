0.  go install ./... to genarate binary in gopath.


1. create user:

orgd tx org create-user name email --from alice


2. delete user:

orgd tx org delete-user userid --from alice   (userid is key of the store (indexes))


3. query all users:

orgd q org users


4.query one user by userid:

orgd q org one-user userid (indexes like 0,1,2...)