import { model, Document, Schema } from "mongoose";

type SocialMedia = "facebook" | "twitter" | "instagram" | "googleBusiness" | "pinterest";

export interface EnterpriseI {
  address: {
    country: string;
    zipcode: string;
    city: string;
    street: string;
  };
  employees: number[];
  media: {
    config: Record<SocialMedia, unknown>;
    posts: unknown[];
    stories: unknown[];
  };
}

export type EnterpriseDoc = EnterpriseI & Document<string>;

const schema = new Schema<EnterpriseDoc>({
  adress: {
    country: { type: String, required: true },
    zipcode: { type: String, required: true },
    city: { type: String, required: true },
    street: { type: String, required: true },
  },
  employees: {
    type: [{ id: { type: Schema.Types.ObjectId, ref: "User" }, role: String }],
    default: [],
  },
  media: {
    config: Schema.Types.Mixed,
    posts: { type: [Schema.Types.Mixed], default: [] },
    stories: { type: [Schema.Types.Mixed], default: [] },
  },
});

export default model<EnterpriseDoc>("enterprise", schema);
