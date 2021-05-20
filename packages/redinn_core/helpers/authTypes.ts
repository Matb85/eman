interface Address {
  country: string;
  zipcode: string;
  city: string;
  street: string;
}

export interface EnterpriseI {
  name: string;
  logo: string;
  address: Address;
}

export interface UserI {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  enterprises: number[];
}
