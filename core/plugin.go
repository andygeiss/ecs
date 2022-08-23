package core

// Plugin is a function which handles a specific kind of functionality
// by using an defaultEntityManager to gain access to the entities.
type Plugin func(em EntityManager) (state int)
