#include <stdio.h>

struct Person  {
  char *name;
  int age;
  double z;
} person;


int main()
{
  struct Person *person1 = &person;
  struct Person *person2 = person1;
  
  person.name = "Susinthiran .S";
  person.age = 45;

  printf("Person sizeof: %d\n", sizeof(person));
  printf("Person address: %x\n", &person);
  printf("Person1 address: %x\n", &person1);
  printf("Person1 follow pointer: %x\n", (*person1));
  printf("Person2 address: %x\n", &person2);
  printf("Person2 follow pointer: %x\n", *person2);
  printf("Person2 name: %s\n", person2->name);

  printf("Person1 follow pointer: %x\n", (*person1));

  
  printf("Name: %s\n", person.name);
  printf("Age: %d\n", person.age);

  printf("Name through pointer: %s\n", person1->name);
  printf("Age through pointer: %d\n", person1->age);

  printf("Name through pointer#2: %s\n", (*person1).name);
  printf("Age through pointer#2: %d", (*person1).age);

}

  
