import bcrypt from "bcrypt";

export async function comparePasswords(plainPassword: string, hashedPassword: string): Promise<boolean> {
  return await bcrypt.compare(plainPassword, hashedPassword);
}

export async function GeneratePasswordHash(password: string): Promise<string> {
  return await bcrypt.hash(password, 12);
}

export default {
  comparePasswords,
  GeneratePasswordHash,
};
