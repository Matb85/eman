import { model, Document, Schema } from "mongoose";

type SocialMedia = "facebook" | "twitter" | "instagram" | "googleBusiness" | "pinterest";

interface employee {
  ref: string;
  role: string;
}

export interface EnterpriseI {
  name: string;
  logo: string;
  address: {
    country: string;
    zipcode: string;
    city: string;
    street: string;
  };
  employees: employee[];
  media: {
    config: Record<SocialMedia, unknown>;
    posts: unknown[];
    stories: unknown[];
  };
}

export type EnterpriseDoc = EnterpriseI & Document;

const schema = new Schema<EnterpriseDoc>({
  name: { type: String, required: true },
  address: {
    country: { type: String, required: true },
    zipcode: { type: String, required: true },
    city: { type: String, required: true },
    street: { type: String, required: true },
  },
  employees: [{ ref: { type: Schema.Types.ObjectId, ref: "User" }, role: String }],
  media: {
    config: Schema.Types.Mixed,
    posts: { type: [Schema.Types.Mixed], default: [] },
    stories: { type: [Schema.Types.Mixed], default: [] },
  },
});

export default model<EnterpriseDoc>("Enterprise", schema);
