"use client";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "@/components/ui/navigation-menu";
import { UserContext } from "@/state/user";
import { useContext } from "react";

const destinations: { title: string; href: string; description: string }[] = [
  {
    title: "Home",
    href: "/",
    description: "Home site of application",
  },
  {
    title: "Home",
    href: "/",
    description: "Home site of application",
  },
];

export default function Navigation() {
  const user = useContext(UserContext);

  return (
    <UserContext.Provider value={user}>
      <NavigationMenu className="w-screen max-w-none m-0 p-0 justify-normal">
        <div className="w-9/12 flex bg-red-600">
          {destinations.map((destination) => {
            return (
              <NavigationMenuList key={destination.title}>
                <NavigationMenuItem className="m-2 p-2">
                  <NavigationMenuLink href={destination.href}>
                    {destination.title}
                  </NavigationMenuLink>
                </NavigationMenuItem>
              </NavigationMenuList>
            );
          })}
        </div>
        <div className="m-auto bg-red-300 w-3/12">
          <div className="border rounded-xl border-black h-12 w-max">
            <p>Pwoer of m</p>
          </div>
        </div>
      </NavigationMenu>
    </UserContext.Provider>
  );
}
