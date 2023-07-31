"use client";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuIndicator,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  NavigationMenuViewport,
} from "@/components/ui/navigation-menu";

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
  return (
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
  );
}
